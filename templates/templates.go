package templates

import "html/template"

var StylesCSS template.CSS
var AlpineJS template.JS
var AppJS template.JS

var AuthPage *template.Template
var HomePage *template.Template

type Data struct {
	StylesCSS template.CSS
	AlpineJS  template.JS
	AppJS     template.JS

	PageTitle string
	Payload   map[string]any
}

func PageData() *Data {
	return &Data{
		StylesCSS: StylesCSS,
		AlpineJS:  AlpineJS,
		AppJS:     AppJS,
	}
}
