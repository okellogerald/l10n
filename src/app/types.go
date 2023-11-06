package app

type Locale = string

type Content map[string]interface{}

type PlaceHolder struct {
	Name string
	Type string
}

type Method struct {
	Name         string
	Description  string
	PlaceHolders []PlaceHolder
	Translations []string
}

type MethodGroup struct {
	Identifier string
	Methods    []Method
}

type MethodsData struct {
	Methods      []Method
	MethodGroups []MethodGroup
}
