package library

import (
	"embed"
	"text/template"
)



func Mytemplate(fs embed.FS, name string) *template.Template {
	return template.Must(template.New("").ParseFS(fs, name))
}
