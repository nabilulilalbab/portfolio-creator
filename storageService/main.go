package main

import (
	"embed"
	"fmt"
	"net/http"
	"text/template"

	"github.com/nabilulilalbab/library"
)

//go:embed static
var staticFS embed.FS

var templates = template.Must(template.ParseGlob("templates/*.html"))

// var Mytemplate = template.Must(template.New("").ParseFS(staticFS,"static"))

func indexHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	err := templates.ExecuteTemplate(w, "index.html", map[string]any{
		"Title": "My Home Page",
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed rendering index: %w", err)
	}
	return http.StatusOK, nil
}




func errorSimulasi(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.URL.Query().Get("error") == "true" {
		return http.StatusInternalServerError, fmt.Errorf("simulated internal error")
	}

	err := templates.ExecuteTemplate(w, "index.html", map[string]any{
		"Title": "My Home Page",
	})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("failed rendering index: %w", err)
	}
	return http.StatusOK, nil
}




func main()  {
 	mux := http.NewServeMux()
	fileserver := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/",library.FileServerWithLog(fileserver))
	mux.HandleFunc("/", library.AppHandler(indexHandler).CreateHandler(library.ErrorHandlerOptions{
		Templates: templates,
	}))

		mux.HandleFunc("/error", library.AppHandler(errorSimulasi).CreateHandler(library.ErrorHandlerOptions{
		Templates: templates,
	}))

	panic(http.ListenAndServe(":8080", mux))
	
}
