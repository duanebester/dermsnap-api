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
	services         *services.Services
	appleVerifier    *oidc.IDTokenVerifier
	appleConfig      oauth2.Config
	doximityVerifier *oidc.IDTokenVerifier
	doximityConfig   oauth2.Config
	googleVerifier   *oidc.IDTokenVerifier
	googleConfig     oauth2.Config
}

func CreateAppleConfig(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("APPLE_CLIENT_ID")
	clientSecret := os.Getenv("APPLE_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("APPLE_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{ClientID: clientID}
	verifier := provider.Verifier(oidcConfig)

	redirectUrl := "https://api.dermsnap.io/oauth2/apple/callback"
	if appEnv == "" || appEnv == "local" {
		redirectUrl = "http://localhost:3000/oauth2/apple/callback"
	}
	if appEnv == "development" {
		redirectUrl = "https://api-dev.dermsnap.io/oauth2/apple/callback"
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   provider.Endpoint().AuthURL,
			TokenURL:  provider.Endpoint().TokenURL,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: redirectUrl,
		Scopes:      []string{oidc.ScopeOpenID, "name", "email"},
	}
	return config, verifier
}

func CreateDoximityConfig(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
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
	return config, verifier
}

func CreateGoogleConfig(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("GOOGLE_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	oidcConfig := &oidc.Config{ClientID: clientID}
	verifier := provider.Verifier(oidcConfig)

	redirectUrl := "https://api.dermsnap.io/oauth2/google/callback"
	if appEnv == "" || appEnv == "local" {
		redirectUrl = "http://localhost:3000/oauth2/google/callback"
	}
	if appEnv == "development" {
		redirectUrl = "https://api-dev.dermsnap.io/oauth2/google/callback"
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectUrl,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	return config, verifier
}

func NewApi(s *services.Services) API {
	ctx := context.Background()
	appleConfig, appleVerifier := CreateAppleConfig(ctx)
	doximityConfig, doximityVerifier := CreateDoximityConfig(ctx)
	googleConfig, googleVerifier := CreateGoogleConfig(ctx)

	return API{
		services:         s,
		appleVerifier:    appleVerifier,
		appleConfig:      appleConfig,
		doximityVerifier: doximityVerifier,
		doximityConfig:   doximityConfig,
		googleVerifier:   googleVerifier,
		googleConfig:     googleConfig,
	}
}
