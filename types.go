package main

type Empty struct{}

type Ast struct {
	Declarations []Declaration `json:"declarations"`
	End          int64         `json:"end"`
	Exports      []interface{} `json:"exports"`
	FilePath     string        `json:"filePath"`
	Imports      []interface{} `json:"imports"`
	Resources    []Resource    `json:"resources"`
	RootPath     string        `json:"rootPath"`
	Start        int64         `json:"start"`
	Usages       []string      `json:"usages"`
}

type Resource struct {
	Declarations []Declaration `json:"declarations"`
	End          int64         `json:"end"`
	Exports      []interface{} `json:"exports"`
	Imports      []interface{} `json:"imports"`
	Name         string        `json:"name"`
	Resources    []interface{} `json:"resources"`
	Start        int64         `json:"start"`
	Usages       []string      `json:"usages"`
}

type Declaration struct {
	Accessors  []interface{} `json:"accessors"`
	End        int64         `json:"end"`
	IsAsync    bool          `json:"isAsync"`
	IsConst    bool          `json:"isConst"`
	IsExported bool          `json:"isExported"`
	Methods    []struct {
		End        int64  `json:"end"`
		IsAbstract bool   `json:"isAbstract"`
		IsAsync    bool   `json:"isAsync"`
		IsOptional bool   `json:"isOptional"`
		IsStatic   bool   `json:"isStatic"`
		Name       string `json:"name"`
		Parameters []struct {
			End   int64  `json:"end"`
			Name  string `json:"name"`
			Start int64  `json:"start"`
			Type  string `json:"type"`
		} `json:"parameters"`
		Start      int64         `json:"start"`
		Type       string        `json:"type"`
		Variables  []interface{} `json:"variables"`
		Visibility int64         `json:"visibility"`
	} `json:"methods"`
	Name       string `json:"name"`
	Parameters []struct {
		End   int64  `json:"end"`
		Name  string `json:"name"`
		Start int64  `json:"start"`
		Type  string `json:"type"`
	} `json:"parameters"`
	Properties []struct {
		End        int64  `json:"end"`
		IsOptional bool   `json:"isOptional"`
		IsStatic   bool   `json:"isStatic"`
		Name       string `json:"name"`
		Start      int64  `json:"start"`
		Type       string `json:"type"`
		Visibility int64  `json:"visibility"`
	} `json:"properties"`
	Start          int64         `json:"start"`
	Type           string        `json:"type"`
	TypeParameters []string      `json:"typeParameters"`
	Variables      []interface{} `json:"variables"`
}
