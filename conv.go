package main

var componentsAdded = make(map[string]string)
var builtins = map[string]string{
	"string": "string",
	"number": "int", // Fix this somehow, may be innacurate
}

func addComponent(name string) string {
	val, exists := componentsAdded[name]
	if exists {
		return val
	}
}
