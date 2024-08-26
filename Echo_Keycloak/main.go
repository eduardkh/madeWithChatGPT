package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lestrrat-go/jwx/jwk"
	"golang.org/x/oauth2"
)

var (
	keycloakConfig = &oauth2.Config{
		ClientID:     "my-go-app",
		ClientSecret: "AK7FlsZHEUENXvW4nURRpBAW8kUgPSuI", // Replace with your actual secret
		RedirectURL:  "http://192.168.1.165:1323/callback",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://192.168.1.165:8080/realms/myrealm/protocol/openid-connect/auth",
			TokenURL: "http://192.168.1.165:8080/realms/myrealm/protocol/openid-connect/token",
		},
		Scopes: []string{"openid", "profile"},
	}
	keycloakCertURL = "http://192.168.1.165:8080/realms/myrealm/protocol/openid-connect/certs"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Public route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to the public page!")
	})

	// Protected route with role-based authorization check
	e.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are viewing an admin page!")
	}, isAuthenticated, isAdmin)

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "You are viewing a user page!")
	}, isAuthenticated, isUser)

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
		return c.Redirect(http.StatusTemporaryRedirect, "/user")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// isAuthenticated middleware to check if the user is authenticated
func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve the token from the cookie
		cookie, err := c.Cookie("access_token")
		if err != nil || cookie.Value == "" {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		// Parse and verify the JWT token
		token, err := jwt.Parse(cookie.Value, fetchKey)
		if err != nil || !token.Valid {
			return c.String(http.StatusUnauthorized, "Invalid token: "+err.Error())
		}

		// Store the token claims for future use
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.String(http.StatusUnauthorized, "Failed to parse claims")
		}
		c.Set("claims", claims)

		return next(c)
	}
}

// isAdmin middleware to check if the user has the 'admin' role
func isAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("claims").(jwt.MapClaims)
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		// Check if the user has the admin role in the resource_access.my-go-app.roles
		resourceAccess, ok := claims["resource_access"].(map[string]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		clientAccess, ok := resourceAccess["my-go-app"].(map[string]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		roles, ok := clientAccess["roles"].([]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		for _, role := range roles {
			if role == "admin" {
				return next(c)
			}
		}

		return c.String(http.StatusForbidden, "You do not have permission to access this page")
	}
}

// isUser middleware to check if the user has the 'user' role
func isUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, ok := c.Get("claims").(jwt.MapClaims)
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		// Check if the user has the user role in the resource_access.my-go-app.roles
		resourceAccess, ok := claims["resource_access"].(map[string]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		clientAccess, ok := resourceAccess["my-go-app"].(map[string]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		roles, ok := clientAccess["roles"].([]interface{})
		if !ok {
			return c.String(http.StatusForbidden, "You do not have permission to access this page")
		}

		for _, role := range roles {
			if role == "user" {
				return next(c)
			}
		}

		return c.String(http.StatusForbidden, "You do not have permission to access this page")
	}
}

// fetchKey fetches the public key for verifying the JWT from Keycloak's JWKS

func fetchKey(token *jwt.Token) (interface{}, error) {
	// Fetch JWKS from Keycloak
	resp, err := http.Get(keycloakCertURL)
	if err != nil {
		return nil, errors.New("Unable to fetch JWKS")
	}
	defer resp.Body.Close()

	// Read the response body into a byte slice
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Failed to read JWKS response body")
	}

	// Parse the JWKS JSON response
	keySet, err := jwk.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWKS: %v", err)
	}

	// Extract the public key by the key ID (kid)
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("No key ID found in token header")
	}

	// Look for the matching key in the set
	key, found := keySet.LookupKeyID(keyID)
	if !found {
		return nil, errors.New("Key not found")
	}

	var raw interface{}
	err = key.Raw(&raw)
	if err != nil {
		return nil, errors.New("Failed to get the raw key")
	}

	return raw, nil
}
