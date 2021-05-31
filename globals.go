package main

import "fmt"

func addGlobals() {
	for _, vr := range vars {
		addComponent(vr.Type)
	}
	for _, k := range types {
		addComponent(k.Type)
	}

	out.WriteString("\n")
	for _, vr := range vars {
		fmt.Fprintf(out, "__global JS.%s %s\n", vr.Name, getKind(vr.Type))
	}

	out.WriteString("\n")
	for _, k := range types {
		fmt.Fprintf(out, "type JS.%s = %s\n", k.Name, getKind(k.Type))
	}
}
