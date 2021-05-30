package main

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed lib.dom.d.ts
var tsCode string

var decls = make(map[string]Interface)
var out = &strings.Builder{}

func main() {
	decls = parseCode(tsCode)

	addComponent("HTMLElement")

	// Save to file
	outFile, err := os.Create("vjs.js.v")
	if err != nil {
		panic(err)
	}

	_, err = outFile.Write([]byte(out.String()))
	if err != nil {
		panic(err)
	}
}
