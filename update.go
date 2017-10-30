// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	log.SetFlags(0)

	// Ensure working directory is right.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.HasSuffix(dir, "/src/github.com/vecty/blink-idl") {
		log.Fatal("error: expected working directory $GOPATH/src/github.com/vecty/blink-idl")
	}

	// Check if the blink repository is cloned or not.
	repoDir := filepath.Join(dir, "blink")
	fi, err := os.Stat(repoDir)
	if os.IsNotExist(err) || err == nil && fi.IsDir() == false {
		log.Fatal(`error: expected repository to be cloned in 'blink' subdirectory

If this is your first time running 'go run update.go', you will need to first clone
the repository. To do this, please run:

$ git clone https://chromium.googlesource.com/chromium/blink

Warning: The above command will download 5 GB of data, and may take upwards of an
hour. Be prepared to make some coffee! =)
`)
	}

	// Pull repository changes.
	log.Println("- Fetching latest changes for ./blink repository")
	execCmd(repoDir, "git", "pull")

	log.Println("- Removing existing ./idl")
	if err := os.RemoveAll("idl"); err != nil {
		log.Fatal(err)
	}

	// Copy IDL files out of the repository.
	log.Println("- Copying IDL files to ./idl")
	count := 0
	err = filepath.Walk(repoDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only interested in IDL files.
		if filepath.Ext(path) != ".idl" {
			return nil
		}

		// Determine path relative to ./blink
		rel, err := filepath.Rel(repoDir, path)
		if err != nil {
			return err
		}

		// Ensure all directories exist.
		newPath := filepath.Join("idl", rel)
		if err := os.MkdirAll(filepath.Dir(newPath), 0777); err != nil {
			return err
		}

		// Copy the file.
		data, err := ioutil.ReadFile(filepath.Join(repoDir, rel))
		if err != nil {
			fmt.Println(rel)
			return err
		}
		err = ioutil.WriteFile(newPath, data, 0777)
		if err != nil {
			return err
		}
		count++
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("- Copied", count, "IDL files to ./idl")

	// Update revision file.
	log.Println("- Updating ./idl/REVISION file")
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = repoDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("idl/REVISION", out, 0777)
	if err != nil {
		log.Fatal(err)
	}

	// Generate VFS.
	execCmd("", "vfsgendev", "-source", `"github.com/vecty/blink-idl".Assets`)

	// Commit the changes.
	log.Println("- Committing changes")
	execCmd("", "git", "add", "./idl", "./idl/REVISION", "./assets_vfsdata.go")
	execCmd("", "git", "commit", "-m", "Update IDL files to blink@"+string(out)[:6])
	log.Println("\nSuccess! You can now `git push`")
}

func execCmd(dir, cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}
}
