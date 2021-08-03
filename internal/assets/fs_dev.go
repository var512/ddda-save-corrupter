// +build dev

package assets

import (
	"net/http"
)

var WebDir = "__invalid__"

var Assets http.FileSystem = http.Dir(WebDir)
