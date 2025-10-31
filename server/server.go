package server

import (
	"embed"
	"html/template"
	"log/slog"
	"maps"
	"net/http"
)

type (
	routes map[string]http.HandlerFunc

	errorType struct {
		Title string
		Msg   string
	}

	jsonApi struct {
		Login bool   `json:"login"`
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

	_dir    = "pages/"
	_layout = _dir + "layout.html"

	_components = "components/*.html"

	_indexTempl *template.Template
	_adminTempl *template.Template
	_errorTempl *template.Template

	_componentsTempl *template.Template
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

func (_routes routes) getCommonRoutes() routes {
	_commonRoutes := routes{
		"/assets/":       assetsWithType,
		"POST /c/{name}": componentsPage,
		"/404/":          pageNotFound,
		"/error/":        errorPage,
	}
	maps.Copy(_routes, _commonRoutes)
	return _routes
}

func (_routes routes) createRoutes() *http.ServeMux {
	_mux := http.NewServeMux()
	for route, handler := range _routes {
		_mux.HandleFunc(route, handler)
	}
	return _mux
}

func New() *http.ServeMux {
	return routes{
		"/": indexPage,
	}.getCommonRoutes().createRoutes()
}

func NewAdmin() *http.ServeMux {
	return routes{
		"/": adminPage,
	}.getCommonRoutes().createRoutes()
}
