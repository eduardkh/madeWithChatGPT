package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer defines a struct that implements the Renderer interface for Echo.
// This allows for custom rendering of HTML templates.
type TemplateRenderer struct {
	templates *template.Template // Holds the compiled templates
}

// Render is a method that Echo calls to render HTML templates.
// It writes the rendered template output to `w`.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// ExecuteTemplate finds the template named `name` and executes it,
	// writing the generated HTML to `w`.
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New() // Initialize a new Echo instance

	// Initialize a TemplateRenderer and load all templates matching the pattern "views/*.gohtml".
	// `template.Must` is a helper that panics if an error occurs, simplifying error handling during initialization.
	t := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.gohtml")),
	}
	// Set the custom renderer for the Echo instance.
	// This tells Echo to use our TemplateRenderer to render HTML responses.
	e.Renderer = t

	// Define a struct to hold data that will be passed to the template.
	type PageData struct {
		Name     string   // The user's name
		LoopData []string // Data to be looped over in the template
	}

	// Create an instance of PageData with some example data.
	pageData := PageData{
		Name:     "Echo",
		LoopData: []string{"One", "Two", "Three"},
	}

	// Define a route for the root URL ("/").
	// When this route is hit, Echo will call the function defined below.
	e.GET("/", func(c echo.Context) error {
		// Render the "index.gohtml" template, passing in `pageData` as the data context.
		return c.Render(http.StatusOK, "index.gohtml", pageData)
	})

	// Start the server on port 1323.
	// `e.Logger.Fatal` logs a fatal error if the server cannot start, causing the application to exit.
	e.Logger.Fatal(e.Start(":1323"))
}
