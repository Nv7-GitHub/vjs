package main

import "strings"

var kindMap = map[string]string{
	"number":  "int",
	"string":  "string",
	"null":    "nil",
	"void":    "nil",
	"any":     "any",
	"boolean": "bool",
	"int":     "int",
}

func init() {
	for _, v := range kindMap {
		added[v] = Empty{}
	}
}

func getKind(kind string) string {
	kind = strings.ReplaceAll(kind, " & ", " | ")

	if strings.Contains(kind, "<") {
		return getKind(kind[:strings.Index(kind, "<")])
	}

	if strings.Contains(kind, "[]") {
		return "[]" + getKind(strings.ReplaceAll(kind, "[]", ""))
	}

	if strings.Contains(kind, "]:") {
		return getKind(strings.Split(kind, "]:")[0])
	}

	if strings.Contains(kind, "=>") {
		return getKind(strings.TrimSpace(strings.Split(strings.ReplaceAll(kind, ")", ""), "=>")[1]))
	}

	if strings.Contains(kind, "|") {
		return getKind(strings.Split(kind, " | ")[0])
	}

	newKind, exists := kindMap[kind]
	if exists {
		return newKind
	}

	return "JS." + kind
}
