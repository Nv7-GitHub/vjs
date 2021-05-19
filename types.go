package main

type Empty struct{}

type Interface struct {
	Methods    []Method
	Properties []Property
	Implements []string
	Name       string
}

type Method struct {
	Parameters []Parameter
	ReturnType string
	Name       string
}

type Parameter struct {
	Name string
	Type string
}

type Property struct {
	IsReadonly bool
	Name       string
	Type       string
}
