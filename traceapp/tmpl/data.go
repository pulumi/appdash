//go:build dev
// +build dev

package tmpl

import (
	"go/build"
	"log"
	"net/http"
)

func importPathToDir(importPath string) string {
	p, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

// Data is a virtual filesystem that contains template data used by Appdash.
var Data = http.Dir(importPathToDir("github.com/pulumi/appdash/traceapp/tmpl/data"))
