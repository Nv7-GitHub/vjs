package main

import "strings"

var kindMap = map[string]string{
	"number":  "JS.Number",
	"string":  "JS.String",
	"null":    "nil",
	"void":    "nil",
	"any":     "voidptr",
	"boolean": "JS.Boolean",
	"int":     "JS.Number",
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
