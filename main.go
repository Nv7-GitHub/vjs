package main

import (
	_ "embed"
	"encoding/json"
	"os"
)

//go:embed ast.json
var astRaw []byte

var decls = make(map[string]Declaration)

func main() {
	var ast Ast
	err := json.Unmarshal(astRaw, &ast)
	if err != nil {
		panic(err)
	}

	for _, decl := range ast.Declarations {
		_, exists := decls[decl.Name]
		if !exists {
			decls[decl.Name] = decl
		}
	}

	// Add HTMLElement
	addComponent("HTMLElement")

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
