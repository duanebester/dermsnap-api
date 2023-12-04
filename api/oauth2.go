package api

import (
	"dermsnap/models"
	"dermsnap/utils"
	"errors"
	"fmt"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/oauth2"
)

func (a API) HandleLoginWithApple(c *fiber.Ctx) error {
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
		Name:     "apple-state",
		Value:    state,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "apple-nonce",
		Value:    nonce,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "apple-verifier",
		Value:    verifier,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})

	redirectUrl := a.services.ProviderService.AppleConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oidc.Nonce(nonce),
		oauth2.S256ChallengeOption(verifier),
		oauth2.SetAuthURLParam("response_mode", "form_post"),
	)
	return c.Redirect(redirectUrl, fiber.StatusFound)
}

type AppleOAuth2CallbackBody struct {
	Code             string `json:"code"`
	State            string `json:"state"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func (a API) HandleAppleOAuth2Callback(c *fiber.Ctx) error {
	ctx := c.Context()

	var respBody AppleOAuth2CallbackBody
	err := c.BodyParser(&respBody)
	if err != nil {
		return err
	}

	errorMessage := respBody.Error
	errorDescription := respBody.ErrorDescription
	if errorMessage != "" {
		log.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
		return fmt.Errorf("oauth2 error: %s - %s", errorMessage, errorDescription)
	}

	storedState := c.Cookies("apple-state")
	if storedState == "" || respBody.State != storedState {
		return errors.New("state did not match")
	}

	codeVerifier := c.Cookies("apple-verifier")
	if codeVerifier == "" {
		return errors.New("code verifier is missing")
	}

	// remove state from cookie
	c.ClearCookie("apple-state")
	c.ClearCookie("apple-nonce")
	c.ClearCookie("apple-verifier")

	oauth2Token, err := a.services.ProviderService.AppleConfig.Exchange(ctx, respBody.Code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		return err
	}
	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("missing token")
	}

	// Parse and verify ID Token payload.
	idToken, err := a.services.ProviderService.AppleVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	// Extract custom claims
	var claims struct {
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
	}
	if err := idToken.Claims(&claims); err != nil {
		log.Errorf("failed to parse claims: %s", err)
		return err
	}

	log.Infof("claims: %+v", claims)

	_, err = a.services.UserService.CreateUser(claims.Email, models.Client, models.Apple)
	if err != nil {
		return err
	}

	token, err := a.services.AuthService.GenerateToken(claims.Email, models.Client, models.Apple)
	if err != nil {
		return err
	}

	// Set token in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "dermsnap-auth",
		Value:    token,
		Secure:   true,
		Domain:   "dermsnap.io",
		MaxAge:   int(time.Hour.Seconds() * 72),
		SameSite: "None",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 72),
	})

	return c.Redirect("https://api-dev.dermsnap.io/", fiber.StatusFound)
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
		Name:     "doximity-state",
		Value:    state,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "doximity-nonce",
		Value:    nonce,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "doximity-verifier",
		Value:    verifier,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})

	redirectUrl := a.services.ProviderService.DoximityConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oidc.Nonce(nonce),
		oauth2.S256ChallengeOption(verifier),
	)
	return c.Redirect(redirectUrl, fiber.StatusFound)
}

func (a API) HandleDoximityOAuth2Callback(c *fiber.Ctx) error {
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

	storedState := c.Cookies("doximity-state")
	if storedState == "" || state != storedState {
		return errors.New("state did not match")
	}

	codeVerifier := c.Cookies("doximity-verifier")
	if codeVerifier == "" {
		return errors.New("code verifier is missing")
	}

	// remove state from cookie
	c.ClearCookie("doximity-state")
	c.ClearCookie("doximity-nonce")
	c.ClearCookie("doximity-verifier")

	oauth2Token, err := a.services.ProviderService.DoximityConfig.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		return err
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("missing token")
	}

	// Parse and verify ID Token payload.
	idToken, err := a.services.ProviderService.DoximityVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

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

	doctor, err := a.services.UserService.CreateUser(idToken.Subject, models.Doctor, models.Doximity)
	if err != nil {
		return err
	}

	_, err = a.services.UserService.CreateDoctorInfo(doctor.ID, claims.Specialty, claims.Credentials)
	if err != nil {
		return err
	}

	token, err := a.services.AuthService.GenerateToken(idToken.Subject, models.Doctor, models.Doximity)
	if err != nil {
		return err
	}

	// Set token in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "dermsnap-auth",
		Value:    token,
		Secure:   true,
		Domain:   "dermsnap.io",
		MaxAge:   int(time.Hour.Seconds() * 72),
		SameSite: "None",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 72),
	})

	return c.Redirect("https://api-dev.dermsnap.io/", fiber.StatusFound)
}

func (a API) HandleLoginWithGoogle(c *fiber.Ctx) error {
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
		Name:     "google-state",
		Value:    state,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "google-nonce",
		Value:    nonce,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "google-verifier",
		Value:    verifier,
		Secure:   true,
		Domain:   "dermsnap.io",
		HTTPOnly: true,
		MaxAge:   int(time.Hour.Seconds()),
		SameSite: "None",
	})

	redirectUrl := a.services.ProviderService.GoogleConfig.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline,
		oidc.Nonce(nonce),
		oauth2.S256ChallengeOption(verifier),
	)
	return c.Redirect(redirectUrl, fiber.StatusFound)
}

func (a API) HandleGoogleOAuth2Callback(c *fiber.Ctx) error {
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

	storedState := c.Cookies("google-state")
	if storedState == "" || state != storedState {
		return errors.New("state did not match")
	}

	codeVerifier := c.Cookies("google-verifier")
	if codeVerifier == "" {
		return errors.New("code verifier is missing")
	}

	// remove state from cookie
	c.ClearCookie("google-state")
	c.ClearCookie("google-nonce")
	c.ClearCookie("google-verifier")

	oauth2Token, err := a.services.ProviderService.GoogleConfig.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		return err
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return errors.New("missing token")
	}

	// Parse and verify ID Token payload.
	idToken, err := a.services.ProviderService.GoogleVerifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	// Extract custom claims
	var claims struct {
		Email      string `json:"email"`
		Name       string `json:"name"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
		Picture    string `json:"picture"`
	}
	if err := idToken.Claims(&claims); err != nil {
		log.Errorf("failed to parse claims: %s", err)
		return err
	}

	log.Infof("claims: %+v", claims)

	found, err := a.services.UserService.GetUserByIdentifier(claims.Email, models.Google)
	if err != nil {
		if err.Error() != "record not found" {
			log.Errorf("failed to find user: %s", err)
			return err
		}
	}

	if found == nil {
		found, err := a.services.UserService.CreateUser(claims.Email, models.Client, models.Google)
		if err != nil {
			log.Errorf("failed to create user: %s", err)
			return err
		}

		log.Infof("user created: %+v", found)
	}

	token, err := a.services.AuthService.GenerateToken(claims.Email, models.Client, models.Google)
	if err != nil {
		log.Errorf("failed to generate token: %s", err)
		return err
	}

	log.Infof("token created: %+v", token)

	// Set token in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "dermsnap-auth",
		Value:    token,
		Secure:   true,
		Domain:   "dermsnap.io",
		MaxAge:   int(time.Hour.Seconds() * 72),
		SameSite: "None",
		HTTPOnly: true,
		Expires:  time.Now().Add(time.Hour * 72),
	})

	return c.Redirect("https://api-dev.dermsnap.io/", fiber.StatusFound)
}
