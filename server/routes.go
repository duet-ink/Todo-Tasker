package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
)

func assetsWithType(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, ".")
	if len(parts) > 1 {
		_type := parts[len(parts)-1]
		mimeType := map[string]string{
			"wasm": "wasm",
			"js":   "javascript",
		}
		if _mime, ok := mimeType[_type]; ok {
			w.Header().Set(
				"Content-Type",
				fmt.Sprintf(
					"application/%s",
					_mime,
				),
			)
		}
	}
	fileServer := http.FileServerFS(_assets)
	fileServer.ServeHTTP(w, r)
}

func componentsPage(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if err := _componentsTempl.ExecuteTemplate(w, name, nil); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/404", http.StatusFound)
	}
	if err := _indexTempl.Execute(w, nil); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}

func adminPage(w http.ResponseWriter, r *http.Request) {
	if err := _adminTempl.Execute(w, nil); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/404", http.StatusFound)
	}
}

// need 400 to 510 error handling

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	if err := _errorTempl.Execute(w, errorType{
		Title: "404",
		Msg:   "Page Not found!",
	}); err != nil {
		slog.Error(err.Error())
		http.Redirect(w, r, "/error", http.StatusFound)
	}
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	if err := _errorTempl.Execute(w, errorType{
		Title: "500",
		Msg:   "Oops! somthing went Wrong!",
	}); err != nil {
		http.Error(w, "Internal error!", http.StatusInternalServerError)
		slog.Error(err.Error())
	}
}
