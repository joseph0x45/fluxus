package main

import (
	"embed"
	"fluxus/templates"
	"html/template"
)

func initTemplates(templatesFS embed.FS) {
	templates.AlpineJS = alpineJS
	templates.AppJS = appJS
	templates.StylesCSS = stylesCSS

	templates.AuthPage = template.Must(
		template.ParseFS(
			templatesFS,
			"templates/base.html",
			"templates/auth.html",
		),
	)
	templates.HomePage = template.Must(
		template.ParseFS(
			templatesFS,
			"templates/base.html",
			"templates/home.html",
		),
	)
}
