package main

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed lib.dom.d.ts
var tsCode string

var decls = make(map[string]Interface)
var vars = make(map[string]Variable)
var types = make(map[string]Type)
var out = &strings.Builder{}

func main() {
	decls = parseCode(tsCode)

	addGlobals()
	for comp := range decls {
		addComponent(comp)
	}

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
