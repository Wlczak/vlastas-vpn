package templates

import (
	"embed"
	"text/template"
)

//go:embed *.tmpl
var templateFS embed.FS

func ParseTemplate(templateName string) template.Template {
	tmpl, err := template.New(templateName).ParseFS(templateFS, templateName)

	if err != nil {
		panic(err)
	}
	return *tmpl
}
