package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed lib.dom.d.ts
var tsCode string

var decls = make(map[string]Interface)
var out string

func main() {
	decls = parseCode(tsCode)

	d := decls["HTMLElement"]
	fmt.Println(d.Name, d.Implements)
	for _, method := range d.Methods {
		fmt.Println("    ", method.Name, method.ReturnType)
		for _, param := range method.Parameters {
			fmt.Println("    ", "    ", param.Name, param.Type)
		}
	}
	for _, prop := range d.Properties {
		fmt.Println("    ", prop.IsReadonly, prop.Name, prop.Type)
	}

	// Save to file
	outFile, err := os.Create("vjs.js.v")
	if err != nil {
		panic(err)
	}

	_, err = outFile.Write([]byte(out))
	if err != nil {
		panic(err)
	}
}
