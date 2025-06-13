package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer satisfies echo.Renderer
// ------------------------------------------------
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

// Panel handlers -----------------------------------------------------------
func Index(c echo.Context) error      { return c.Render(http.StatusOK, "index.html", nil) }
func PingPanel(c echo.Context) error  { return c.Render(http.StatusOK, "ping.html", nil) }
func DNSPanel(c echo.Context) error   { return c.Render(http.StatusOK, "dns.html", nil) }
func TracePanel(c echo.Context) error { return c.Render(http.StatusOK, "trace.html", nil) }

// runPingWithTimestamp runs ping with -D and parses epoch timestamps
func runPingWithTimestamp(host string) (string, error) {
	cmd := exec.Command("ping", "-4", "-c", "4", "-i", "0.1", "-D", host)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(stdout)
	var buf bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		// parse [epoch.sec] prefix
		if strings.HasPrefix(line, "[") {
			end := strings.Index(line, "]")
			if end > 0 {
				tsStr := line[1:end]
				rest := strings.TrimSpace(line[end+1:])

				if secs, err := strconv.ParseFloat(tsStr, 64); err == nil {
					secPart := int64(secs)
					nsec := int64((secs - float64(secPart)) * 1e9)
					t := time.Unix(secPart, nsec).Format(time.RFC3339)
					buf.WriteString(fmt.Sprintf("%s %s\n", t, rest))
					continue
				}
			}
		}
		// fallback
		buf.WriteString(line + "\n")
	}
	cmd.Wait()
	return buf.String(), nil
}

// runCmdWithTimestamps prefixes each output line with the current timestamp
func runCmdWithTimestamps(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(stdout)
	var buf bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		t := time.Now().Format(time.RFC3339Nano)
		buf.WriteString(fmt.Sprintf("%s %s\n", t, line))
	}
	cmd.Wait()
	return buf.String(), nil
}

// Action handlers ----------------------------------------------------------
func Ping(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runPingWithTimestamp(host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}

func DNS(c echo.Context) error {
	host := c.FormValue("host")
	out, err := runCmdWithTimestamps("dig", "+noall", "+answer", host)
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
	out, err := runCmdWithTimestamps("mtr", "-r", "-c", "10", host)
	code := http.StatusOK
	if err != nil {
		code = http.StatusInternalServerError
	}
	return c.HTML(code, "<pre>"+out+"</pre>")
}
