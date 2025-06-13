package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer satisfies echo.Renderer
type TemplateRenderer struct {
	templates *template.Template
}

func NewRenderer(glob string) *TemplateRenderer {
	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(glob)),
	}
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// --- Panel handlers --------------------------------------------------------

func Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
func PingPanel(c echo.Context) error {
	return c.Render(http.StatusOK, "ping.html", nil)
}
func DNSPanel(c echo.Context) error {
	return c.Render(http.StatusOK, "dns.html", nil)
}
func TracePanel(c echo.Context) error {
	return c.Render(http.StatusOK, "trace.html", nil)
}

// --- Action handlers -------------------------------------------------------

func runCmd(args ...string) (string, error) {
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	return string(out), err
}

func Ping(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runCmd("ping", "-c", "4", host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}

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

func Trace(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runCmd("mtr", "-r", "-c", "10", host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}
