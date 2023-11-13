package app

type Locale = string

type Content map[string]interface{}
type ContentList []map[string]interface{}

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
	// Will be the class name for the generated class
	Identifier string
	// Will be the method name in the app_localization file
	Name       string
	Methods    []Method
}

type MethodsData struct {
	Methods      []Method
	MethodGroups []MethodGroup
}
