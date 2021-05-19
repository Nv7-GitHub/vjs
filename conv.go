package main

import (
	"fmt"
)

var out = ""

var componentsAdded = make(map[string]Empty)
var builtins = map[string]string{
	"string":  "string",
	"number":  "int", // Fix this somehow, may be innacurate
	"boolean": "bool",
	"null":    "nil",
}

func addComponent(name string) {
	fmt.Println(name)
	_, exists := componentsAdded[name]
	if exists {
		fmt.Println(name)
		return
	}

	// Add Dependencies for properties
	decl := decls[name]
	for _, prop := range decl.Properties {
		fixType(prop.Type)
	}

	// Add struct and properties
	out += fmt.Sprintf("interface JS.%s {\n", decl.Name)
	for _, prop := range decl.Properties {
		// " | " Types, need to find which one to select, right now takes non-builtin one
		prop.Type = fixType(prop.Type)

		mut := ""
		if !prop.Isstat
		out += fmt.Sprintf("%s	%s %s\n", prop.Name, prop.Type)
	}
	out += "}\n"

	// Add methods
	for _, m := range decl.Methods {
		me := fmt.Sprintf("fn (JS.%s) JS.%s(", decl.Name, m.Name)
		// Add parameters for method
		for _, param := range m.Parameters {
			// Parameter name checking
			if param.Name == "type" {
				param.Name = "type_"
			}

			// Add dependencies for method
			param.Type = fixType(param.Type)

			me += param.Name + " " + param.Type
		}

		// Add return type for method
		if m.Type != "void" {
			// Add dependency for return type
			m.Type = fixType(m.Type)

			me += " " + m.Type
		}

		out += me + ") {} \n"
	}

	componentsAdded[name] = Empty{}
}
