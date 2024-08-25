package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/oauth2"
)

var (
	keycloakConfig = &oauth2.Config{
		ClientID:     "my-go-app",
		ClientSecret: "vdOwNrPhBMFfVUFIVFnGGwToGq8Xyfuq", // Replace with actual secret
		RedirectURL:  "http://192.168.1.165:1323/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://192.168.1.165:8080/realms/myrealm/protocol/openid-connect/auth",
			TokenURL: "http://192.168.1.165:8080/realms/myrealm/protocol/openid-connect/token",
		},
		Scopes: []string{"openid", "profile"},
	}
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the public page!")
	})

	// Protected route
	e.GET("/protected", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are viewing a protected page!")
	}, isAuthenticated)

	// OAuth2 login
	e.GET("/login", func(c echo.Context) error {
		url := keycloakConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})

	// OAuth2 callback
	e.GET("/callback", func(c echo.Context) error {
		code := c.QueryParam("code")
		state := c.QueryParam("state")

		if code == "" || state == "" {
			return c.String(http.StatusBadRequest, "Missing code or state in callback")
		}

		// Exchange the code for a token
		token, err := keycloakConfig.Exchange(c.Request().Context(), code)
		if err != nil {
			fmt.Printf("Error exchanging token: %v\n", err)
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Error exchanging token: %v", err.Error()))
		}

		// Store the token in a cookie (for demonstration purposes)
		cookie := new(http.Cookie)
		cookie.Name = "access_token"
		cookie.Value = token.AccessToken
		cookie.HttpOnly = true
		cookie.Secure = false // Set to true if you're using HTTPS
		c.SetCookie(cookie)

		fmt.Println("Access token:", token.AccessToken)
		return c.Redirect(http.StatusTemporaryRedirect, "/protected")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the token from the cookie
		cookie, err := c.Cookie("access_token")
		if err != nil || cookie.Value == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		// (Optional) Here, you would typically verify the token's validity with Keycloak

		return next(c)
	}
}
