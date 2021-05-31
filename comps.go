package main

import (
	"fmt"
	"strings"
)

var added = make(map[string]Empty)

func addComponent(comp string) {
	comp = strings.ReplaceAll(comp, "JS.", "")

	_, exists := added[comp]
	if exists {
		return
	}
	_, exists = kindMap[comp]
	if exists {
		return
	}
	added[comp] = Empty{}

	c := decls[comp]
	for _, comp := range c.Implements {
		addComponent(comp)
	}

	for i, field := range c.Properties {
		kind := getKind(field.Type)
		addComponent(kind)
		c.Properties[i].Type = kind
	}

	for _, method := range c.Methods {
		addComponent(getKind(method.ReturnType))
		for _, param := range method.Parameters {
			addComponent(getKind(param.Type))
		}
	}

	fmt.Fprintf(out, "interface %s {\n", getKind(comp))

	for _, val := range c.Implements {
		out.WriteString("  " + val + "\n")
	}

	// Add readonly
	for _, field := range c.Properties {
		addField(field)
	}
	if len(c.Properties) > 0 && len(c.Methods) > 0 {
		out.WriteString("\n")
	}

	for _, method := range c.Methods {
		retKind := getKind(method.ReturnType)

		if retKind == "void" {
			retKind = ""
		} else {
			addComponent(retKind)
			retKind = " " + retKind
		}

		params := &strings.Builder{}
		for i, param := range method.Parameters {
			k := getKind(param.Type)
			addComponent(k)

			if param.Name == "type" {
				param.Name = "typ"
			}

			fmt.Fprintf(params, "%s %s", param.Name, k)

			if i != len(method.Parameters)-1 {
				fmt.Fprintf(params, ", ")
			}
		}

		fmt.Fprintf(out, "  %s(%s)%s\n", method.Name, params.String(), retKind)
	}

	out.WriteString("}\n\n")
}

func addField(field Property) {
	if field.Name[0] == '[' {
		field.Name = field.Name[1:]
	}

	if field.Name == "type" {
		field.Name = "typ"
	}

	fmt.Fprintf(out, "  %s %s\n", strings.ReplaceAll(field.Name, "?", ""), field.Type)
}
