package api

import (
	"context"
	"dermsnap/services"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type API struct {
	services *services.Services
	verifier *oidc.IDTokenVerifier
	config   oauth2.Config
}

func NewApi(s *services.Services) API {
	ctx := context.Background()
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("DOXIMITY_CLIENT_ID")
	clientSecret := os.Getenv("DOXIMITY_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("DOXIMITY_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{ClientID: clientID}
	verifier := provider.Verifier(oidcConfig)

	redirectUrl := "https://api.dermsnap.io/oauth2/doximity/callback"
	if appEnv == "" || appEnv == "local" {
		redirectUrl = "http://localhost:3000/oauth2/doximity/callback"
	}
	if appEnv == "development" {
		redirectUrl = "https://api-dev.dermsnap.io/oauth2/doximity/callback"
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectUrl,
		Scopes:       []string{oidc.ScopeOpenID, "profile:read:email", "profile:read:basic"},
	}

	return API{
		services: s,
		verifier: verifier,
		config:   config,
	}
}
