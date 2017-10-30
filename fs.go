// +build dev

package idl

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("idl")
