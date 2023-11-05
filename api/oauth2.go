package api

import (
	"dermsnap/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/oauth2"
)

func SetCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   int(time.Hour.Seconds()),
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}

func (a API) HandleLoginWithDoximity(c *fiber.Ctx) error {
	state, err := utils.RandString(16)
	if err != nil {
		panic(err)
	}
	nonce, err := utils.RandString(16)
	if err != nil {
		panic(err)
	}

	// use PKCE to protect against CSRF attacks
	verifier := oauth2.GenerateVerifier()

	c.Cookie(&fiber.Cookie{
		Name:     "state",
		Value:    state,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
	})
	c.Cookie(&fiber.Cookie{
		Name:     "nonce",
		Value:    nonce,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
	})
	c.Cookie(&fiber.Cookie{
		Name:     "verifier",
		Value:    verifier,
		Secure:   true,
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
	})

	redirectUrl := a.config.AuthCodeURL(state, oauth2.AccessTypeOffline, oidc.Nonce(nonce), oauth2.S256ChallengeOption(verifier))
	return c.Redirect(redirectUrl, fiber.StatusFound)
}

func (a API) HandleOAuth2Callback(c *fiber.Ctx) error {
	ctx := c.Context()

	errorMessage := c.Query("error")
	errorDescription := c.Query("error_description")
	if errorMessage != "" {
		log.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
		return fmt.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
	}

	code := c.Query("code")
	if code == "" {
		return errors.New("code is missing")
	}

	state := c.Query("state")
	if state == "" {
		return errors.New("state is missing")
	}

	storedState := c.Cookies("state")
	if storedState == "" || state != storedState {
		return errors.New("state did not match")
	}

	codeVerifier := c.Cookies("verifier")
	if codeVerifier == "" {
		return errors.New("code verifier is missing")
	}

	oauth2Token, err := a.config.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		return err
	}

	log.Infof("OAuth2 access_token: %s", oauth2Token.AccessToken)

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("missing token")
	}

	log.Infof("rawIDToken: %s", rawIDToken)

	// Parse and verify ID Token payload.
	idToken, err := a.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	log.Info("ID Token: ", idToken)

	// Extract custom claims
	var claims struct {
		Specialty   string `json:"specialty"`
		Credentials string `json:"credentials"`
	}
	if err := idToken.Claims(&claims); err != nil {
		log.Errorf("failed to parse claims: %s", err)
		return err
	}

	log.Infof("claims: %+v", claims)

	return nil
}
