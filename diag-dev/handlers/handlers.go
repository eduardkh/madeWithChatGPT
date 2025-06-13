package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer renders templates/*.html
type TemplateRenderer struct {
	templates *template.Template
}

// NewRenderer loads all templates matching the glob
func NewRenderer(glob string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(glob)),
	}
}

// Render satisfies echo.Renderer
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Index shows the form
func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func runCmd(args ...string) (string, error) {
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	return string(out), err
}

// Ping handler
func Ping(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runCmd("ping", "-c", "4", host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}

// DNS handler (uses dig)
func DNS(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runCmd("dig", "+noall", "+answer", host)
	code := http.StatusOK
	if err != nil || out == "" {
		code = http.StatusInternalServerError
		if out == "" {
			out = "no DNS answers"
		}
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}

// Trace handler (uses mtr in report mode)
func Trace(c echo.Context) error {
	host := c.FormValue("host")
	// mtr report (-r) for a quick summary, limit to 10 hops
	out, err := runCmd("mtr", "-r", "-c", "10", host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}
