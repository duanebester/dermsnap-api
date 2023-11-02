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
	return c.Redirect(a.config.AuthCodeURL(state, oidc.Nonce(nonce)), fiber.StatusFound)
}

func (a API) HandleOAuth2Callback(c *fiber.Ctx) error {
	ctx := c.Context()

	state := c.Query("state")
	if state == "" {
		return errors.New("state is missing")
	}

	code := c.Query("code")
	if state == "" {
		return errors.New("state is missing")
	}

	storedState := c.Cookies("state")
	if storedState == "" || state != storedState {
		return errors.New("state did not match")
	}

	errorMessage := c.Query("error")
	errorDescription := c.Query("error_description")

	if errorMessage != "" {
		log.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
		return fmt.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
	}

	oauth2Token, err := a.config.Exchange(ctx, code)
	if err != nil {
		return err
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("missing token")
	}

	// Parse and verify ID Token payload.
	idToken, err := a.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	// Extract custom claims
	var claims struct {
		Email    string `json:"email"`
		Verified bool   `json:"email_verified"`
	}
	if err := idToken.Claims(&claims); err != nil {
		return err
	}

	return nil
}
