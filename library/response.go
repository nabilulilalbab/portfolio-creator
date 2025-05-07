package library

import (
	"embed"
	"html/template" // Ganti text/template ke html/template
)

// MyTemplate membuat template baru dari embedded filesystem
func MyTemplate(fs embed.FS, pattern string) *template.Template {
	return template.Must(template.New("").ParseFS(fs, pattern))
}
