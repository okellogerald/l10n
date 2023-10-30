package flutter

type PlaceHolder struct {
	Name string
	Type string
}

type Method struct {
	Name         string
	Description  string
	PlaceHolders []PlaceHolder
}
