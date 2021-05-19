package main

import (
	"fmt"
	"strings"
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
	_, exists := componentsAdded[name]
	if exists {
		fmt.Println(name)
		return
	}

	// Add Dependencies for properties
	decl := decls[name]
	for _, prop := range decl.Properties {
		_, exists := builtins[prop.Type]
		if !exists {
			addComponent(prop.Type)
		}
	}

	// Add struct and properties
	out += fmt.Sprintf("interface JS.%s {\n", name)
	for _, prop := range decl.Properties {
		// " | " Types, need to find which one to select, right now takes non-builtin one
		if strings.Contains(prop.Type, "|") {
			types := strings.Split(prop.Type, "|")
			for i, typ := range types {
				types[i] = strings.TrimSpace(typ)
			}

			hasFound := false
			for _, typ := range types {
				_, exists := builtins[typ]
				if !exists {
					addComponent(typ)
					hasFound = true
					prop.Type = "JS." + typ
				}
			}

			// Haven't found non-builtin type, just use the first one (also needs to be fixed)
			if !hasFound {
				prop.Type = types[0]
			}
		}

		conved, exists := builtins[prop.Type]
		if exists {
			prop.Type = conved
		}

		out += fmt.Sprintf("	%s %s\n", prop.Name, prop.Type)
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
			conved, exists := builtins[param.Type]
			if !exists {
				addComponent(param.Type)
				param.Type = "JS." + param.Type
			} else {
				param.Type = conved
			}

			me += param.Name + " " + param.Type
		}

		// Add return type for method
		if m.Type != "void" {
			// Add dependency for return type
			conved, exists := builtins[m.Type]
			if !exists {
				addComponent(m.Type)
				m.Type = "JS." + m.Type
			} else {
				m.Type = conved
			}

			me += " " + m.Type
		}

		out += me + ") {} \n"
	}
}
