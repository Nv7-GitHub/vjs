package main

import (
	_ "embed"
	"encoding/json"
)

//go:embed ast.json
var astRaw []byte

var decls map[string]Declaration

func main() {
	var ast Ast
	json.Unmarshal(astRaw, &ast)

	for _, decl := range ast.Declarations {
		decls[decl.Name] = decl
	}
}
