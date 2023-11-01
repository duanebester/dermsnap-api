package api

import (
	"context"
	"dermsnap/api/http"
	"dermsnap/api/public"
	"dermsnap/database"
	"dermsnap/middleware"
	"dermsnap/services"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"golang.org/x/oauth2"
)

type API struct {
	services *services.Services
	verifier *oidc.IDTokenVerifier
	config   oauth2.Config
}

func newApi(s *services.Services) API {
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

	redirectUrl := "https://api.dermsnap.io/oauth/doximity/callback"
	if appEnv == "" || appEnv == "local" {
		redirectUrl = "http://localhost:3000/oauth/doximity/callback"
	}
	if appEnv == "development" {
		redirectUrl = "https://dev-api.dermsnap.io/oauth/doximity/callback"
	}

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectUrl,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	return API{
		services: s,
		verifier: verifier,
		config:   config,
	}
}

func NewApp() *fiber.App {
	var appEnv = os.Getenv("APP_ENV")
	var db = database.NewDatabase()
	var services = services.NewServices(db)
	var api = newApi(services)

	var engine = html.New("api/views", ".html")
	if appEnv == "" || appEnv == "local" {
		// Reload the templates on each render, good for development
		engine.Reload(true)
		// Debug will print each template that is parsed, good for debugging
		engine.Debug(true)
	}

	var app = fiber.New(fiber.Config{
		AppName:     "dermsnap",
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(pprof.New())
	app.Get("/metrics", monitor.New())
	app.Static("/", "./assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	var publicHandler = public.NewStrictHandler(api, nil)
	var apiHandler = http.NewStrictHandler(api, nil)

	publicRoute := app.Group("/public")
	apiRoute := app.Group("/api", middleware.Protected())

	public.RegisterHandlers(publicRoute, publicHandler)
	http.RegisterHandlers(apiRoute, apiHandler)

	return app
}
