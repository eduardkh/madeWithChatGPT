package handlers

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

var funcMap = template.FuncMap{
	"previousPage": func(x int) int { return x - 1 },
	"nextPage":     func(x int) int { return x + 1 },
}

func TemplateRenderer() echo.Renderer {
	tmpl := template.New("").Funcs(funcMap)
	tmpl, err := tmpl.ParseGlob("views/*.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err = tmpl.ParseGlob("views/partials/*.html")
	if err != nil {
		log.Fatal(err)
	}
	return &templateRenderer{
		templates: tmpl,
	}
}

type templateRenderer struct {
	templates *template.Template
}

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
