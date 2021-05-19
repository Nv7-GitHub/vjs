package main

import "strings"

func fixType(typ string) string {
	if strings.Contains(typ, "|") {
		types := strings.Split(typ, "|")
		for i, typ := range types {
			types[i] = strings.TrimSpace(typ)
		}

		hasFound := false
		for _, ty := range types {
			_, exists := builtins[ty]
			if !exists {
				addComponent(ty)
				hasFound = true
				typ = "JS." + ty
			}
		}

		// Haven't found non-builtin type, just use the first one (also needs to be fixed)
		if !hasFound {
			typ = types[0]
		}
	}

	conved, exists := builtins[typ]
	if exists {
		typ = conved
	} else {
		addComponent(typ)
	}
	return typ
}
