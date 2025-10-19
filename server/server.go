package server

import (
	"embed"
	"html/template"
	"log/slog"
	"net/http"
)

type (
	routes map[string]http.HandlerFunc

	errorType struct {
		Title string
		Msg   string
	}

	jsonApi struct {
		Login bool `json:"login"`
		Todos []todo `json:"todos"`
	}

	todo struct {
		Title       string `json:"title"`
		Url         string `json:"url"`
		Img         string `json:"img"`
		Alt         string `json:"alt"`
		Description string `json:"description"`
		UserName    string `json:"user_name"`
		CreatedAt   string `json:"created_at"`
	}
)

var (
	//go:embed pages
	_pages embed.FS

	//go:embed assets
	_assets embed.FS

	_dir        = "pages/"
	_layout     = _dir + "layout.html"
	_components = "components/*.html"

	_indexTempl      *template.Template
	_adminTempl      *template.Template
	_errorTempl      *template.Template
	_componentsTempl *template.Template

	_routes = routes{
		"/":       indexPage,
		"/404/":   pageNotFound,
		"/admin/": adminPage,
		"/error/": errorPage,

		"POST /components/{name}": componentsPage,
	}
)

func init() {
	_indexTempl = getTemplate("index.html")
	_adminTempl = getTemplate("admin.html")
	_errorTempl = getTemplate("error.html")
	_componentsTempl = getTemplate(_components)
}

func getTemplate(filename string) *template.Template {
	temp, err := template.ParseFS(_pages, _layout, _dir+filename)
	if err != nil {
		slog.Error(err.Error())
	}
	return temp
}

func New() *http.ServeMux {
	_mux := http.NewServeMux()
	_mux.Handle("/assets/", http.FileServerFS(_assets))
	for _route, _handler := range _routes {
		_mux.HandleFunc(
			_route, _handler,
		)
	}
	return _mux
}
