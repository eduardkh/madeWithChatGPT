package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open-policy-agent/opa/rego"
)

// OPAAuthMiddleware checks requests against OPA policies
func OPAAuthMiddleware(policyContent string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := context.Background()

			// Define input for OPA query
			input := map[string]interface{}{
				"user":   c.Request().Header.Get("User"),
				"path":   c.Request().URL.Path,
				"method": c.Request().Method,
			}

			// Prepare and evaluate the policy query
			query, err := rego.New(
				rego.Query("data.myapp.authz.allow"),
				rego.Module("policy.rego", policyContent),
			).PrepareForEval(ctx)
			if err != nil {
				return err // Or handle more gracefully
			}

			results, err := query.Eval(ctx, rego.EvalInput(input))
			if err != nil || len(results) == 0 || !results[0].Expressions[0].Value.(bool) {
				return echo.NewHTTPError(http.StatusForbidden, "Access denied")
			}

			return next(c)
		}
	}
}

func main() {
	e := echo.New()

	// Load policy content from file
	policyContent, err := ioutil.ReadFile("policy.rego")
	if err != nil {
		log.Fatalf("Failed to load policy: %v", err)
	}

	// Use the OPA authorization middleware
	e.Use(OPAAuthMiddleware(string(policyContent)))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
