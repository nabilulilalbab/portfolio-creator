package handler

import (
	"embed"
	"fmt"
	"html/template" // Ganti text/template ke html/template
	"net/http"
)

// Templates akan diinisialisasi dari embedded filesystem
var Templates *template.Template

// InitTemplates menginisialisasi template dari embedded filesystem
func InitTemplates(fs embed.FS) (*template.Template, error) {
	tmpl, err := template.ParseFS(fs, "templates/*.html")
	if err != nil {
		return nil, err
	}
	Templates = tmpl
	return tmpl, nil
}

func IndexHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	err := Templates.ExecuteTemplate(w, "index.html", map[string]any{
		"Title": "My Home Page",
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed rendering index: %w", err)
	}
	return http.StatusOK, nil
}

func ErrorSimulasi(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.URL.Query().Get("error") == "true" {
		return http.StatusInternalServerError, fmt.Errorf("simulated internal error")
	}
	err := Templates.ExecuteTemplate(w, "index.html", map[string]any{
		"Title": "My Home Page",
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed rendering index: %w", err)
	}
	return http.StatusOK, nil
}
