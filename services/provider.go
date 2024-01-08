package services

import (
	"context"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type ProviderService struct {
	AppleVerifier    *oidc.IDTokenVerifier
	AppleConfig      oauth2.Config
	DoximityVerifier *oidc.IDTokenVerifier
	DoximityConfig   oauth2.Config
	GoogleVerifier   *oidc.IDTokenVerifier
	GoogleConfig     oauth2.Config
}

func NewProviderService() ProviderService {
	ctx := context.Background()
	appleConfig, appleVerifier := CreateAppleProvider(ctx)
	doximityConfig, doximityVerifier := CreateDoximityProvider(ctx)
	googleConfig, googleVerifier := CreateGoogleProvider(ctx)
	return ProviderService{
		AppleVerifier:    appleVerifier,
		AppleConfig:      appleConfig,
		DoximityVerifier: doximityVerifier,
		DoximityConfig:   doximityConfig,
		GoogleVerifier:   googleVerifier,
		GoogleConfig:     googleConfig,
	}
}

func CreateAppleProvider(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("APPLE_CLIENT_ID")
	clientSecret := os.Getenv("APPLE_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("APPLE_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		panic(err)
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

func CreateDoximityProvider(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("DOXIMITY_CLIENT_ID")
	clientSecret := os.Getenv("DOXIMITY_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("DOXIMITY_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		panic(err)
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

func CreateGoogleProvider(ctx context.Context) (oauth2.Config, *oidc.IDTokenVerifier) {
	appEnv := os.Getenv("APP_ENV")
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	providerBaseUrl := os.Getenv("GOOGLE_PROVIDER_BASE_URL")
	provider, err := oidc.NewProvider(ctx, providerBaseUrl)
	if err != nil {
		panic(err)
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
