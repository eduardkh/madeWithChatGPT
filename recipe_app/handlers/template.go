package handlers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

func TemplateRenderer() echo.Renderer {
	return &templateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type templateRenderer struct {
	templates *template.Template
}

func (t *templateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
