package muxer

import (
	"path"
	"strings"
)

func shiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i == 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}
