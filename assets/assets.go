package assets

import (
	_ "embed"
)

//go:embed vendor/styles.css
var StylesCSS string

//go:embed app.js
var AppJS string
