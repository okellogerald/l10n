package app

import "github.com/iancoleman/strcase"

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

func (g Method) GetMethodName() string {
	return strcase.ToLowerCamel(g.Name)
}

// Independent classes will be created out of this struct
type MethodGroup struct {
	// Will be the class name for the generated class
	Identifier string

	Methods []Method
}

func (g MethodGroup) GetGroupClassName() string {
	return strcase.ToCamel(g.Identifier)
}

func (g MethodGroup) GetGroupMethodName() string {
	return strcase.ToLowerCamel(g.Identifier)
}

type MethodsData struct {
	Methods      []Method
	MethodGroups []MethodGroup
}
