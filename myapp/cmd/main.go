package main

import (
	"fmt"
	"log"
	"net/http"

	"myapp/config"
	"myapp/internal/server"
	"myapp/internal/vault"

	"github.com/gorilla/sessions"
	"github.com/hashicorp/vault/api"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create Vault API client
	vaultConfig := &api.Config{
		Address: cfg.VaultAddr,
	}

	vaultClient, err := vault.NewClient(vaultConfig)
	if err != nil {
		log.Fatalf("Failed to create Vault client: %v", err)
	}

	// Authenticate with Vault using AppRole
	token, err := vaultClient.Authenticate(cfg.AppRoleID, cfg.AppSecretID)
	if err != nil {
		log.Fatalf("Failed to authenticate with Vault: %v", err)
	}
	vaultClient.SetToken(token)

	// Set up session store
	store := sessions.NewCookieStore([]byte("secret-key")) // Replace "secret-key" with a more secure key

	// Set up router
	router := server.NewRouter(store, vaultClient)

	// Start server
	fmt.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
