package app

import (
	"dermsnap/api"
	"dermsnap/api/http"
	"dermsnap/api/public"
	"dermsnap/database"
	"dermsnap/middleware"
	"dermsnap/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func NewApp() *fiber.App {
	var appEnv = os.Getenv("APP_ENV")
	var db = database.NewDatabase()
	var services = services.NewServices(db)
	var api = api.NewApi(services)

	var engine = html.New("app/views", ".html")
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
