package app

import (
	"dermsnap/api"
	"dermsnap/api/http"
	"dermsnap/api/public"
	"dermsnap/database"
	"dermsnap/middleware"
	"dermsnap/models"
	"dermsnap/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v5"
)

func HandleSeed(services *services.Services) {
	// Seed DB with admin user
	found, err := services.UserService.GetUserByIdentifier("admin", models.Google)
	if err != nil {
		log.Infof("GetUserByIdentifier Error: %+v", err)
		if err.Error() != "record not found" {
			panic(err)
		}
	}
	if found == nil {
		log.Infof("Seeding database with admin user")
		_, err := services.UserService.CreateUser("admin", models.Admin, models.Google)
		if err != nil {
			panic(err)
		}
	}

	// Get token and log to console
	token, err := services.AuthService.GenerateToken("admin", models.Admin, models.Google)
	if err != nil {
		panic(err)
	}
	log.Infof("Admin Token: %s", token)
}

func NewApp(seedData bool) *fiber.App {
	appEnv := os.Getenv("APP_ENV")
	db := database.NewDatabase()
	services := services.NewServices(db)
	api := api.NewApi(services)

	engine := html.New("app/views", ".html")
	if appEnv == "" || appEnv == "local" {
		// Reload the templates on each render, good for development
		engine.Reload(true)
		// Debug will print each template that is parsed, good for debugging
		engine.Debug(true)

		// Handle Setup
		if seedData {
			HandleSeed(services)
		}
	}

	app := fiber.New(fiber.Config{
		AppName:     "dermsnap",
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(pprof.New())
	app.Get("/metrics", monitor.New())
	app.Static("/assets", "/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		// check if user is authenticated
		token := c.Cookies("dermsnap-auth")

		if token == "" {
			return c.Redirect("/login")
		}

		// decode token
		// check if token is valid
		// if token is valid, render index
		// if token is invalid, redirect to /login
		jwtSecret := os.Getenv("JWT_SECRET")
		t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			return c.Redirect("/login")
		}
		if !t.Valid {
			return c.Redirect("/login")
		}

		claims := t.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(string)

		user, err := services.UserService.GetUserByID(userId)
		if err != nil {
			return c.Redirect("/login")
		}

		return c.Render("index", fiber.Map{
			"User": *user,
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})

	app.Get("/login/apple", api.HandleLoginWithApple)
	app.Post("/oauth2/apple/callback", api.HandleAppleOAuth2Callback)

	app.Get("/login/doximity", api.HandleLoginWithDoximity)
	app.Get("/oauth2/doximity/callback", api.HandleDoximityOAuth2Callback)

	app.Get("/login/google", api.HandleLoginWithGoogle)
	app.Get("/oauth2/google/callback", api.HandleGoogleOAuth2Callback)

	publicRoute := app.Group("/public")
	apiRoute := app.Group("/api", middleware.Protected(), middleware.EnrichUser(services.UserService))

	public.RegisterHandlers(publicRoute, public.NewStrictHandler(api, nil))
	http.RegisterHandlers(apiRoute, http.NewStrictHandler(api, nil))

	return app
}
