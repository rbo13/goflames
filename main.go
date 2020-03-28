package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/compression"
	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/limiter"
	"github.com/gofiber/logger"
	"github.com/gofiber/recover"
	"github.com/joho/godotenv"
	"github.com/rbo13/flame"
)

const (
	MaxTimeout    = 15 // in seconds.
	MaxWait       = 5  // in seconds.
	FALLBACK_PORT = 8000
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env File is missing!")
	}

	PORT := os.Getenv("PORT")
	app := fiber.New()

	app.Settings.TemplateFolder = "./public/views"

	app.Static("/assets", "./public/assets")
	app.Use(
		compression.New(),
		helmet.New(),
		recover.New(recover.Config{
			Handler: func(c *fiber.Ctx, err error) {
				c.SendString(err.Error())
				c.SendStatus(http.StatusInternalServerError)
			},
		}),
		limiter.New(limiter.Config{
			Timeout: MaxTimeout,
			Max:     MaxWait,
		}),
		logger.New(logger.Config{
			Format: "${ip} -> ${protocol}:// ${url} :: ${method}. ${time} -> ${ua}\n\n",
		}),
	)

	// serve index page
	app.Get("/", func(c *fiber.Ctx) {
		c.Render("index.html", nil)
	})

	// handle submit
	app.Post("/", func(c *fiber.Ctx) {
		name := strings.ToLower(c.FormValue("name"))
		partner := strings.ToLower(c.FormValue("partner"))

		pair := flame.Pair(name, partner)
		result := flame.GetResult(pair)

		data := fiber.Map{
			"result":    result,
			"name":      name,
			"partner":   partner,
			"generated": true,
		}

		c.Render("index.html", data)
	})

	// fallback port
	go func() {
		app.Listen(FALLBACK_PORT)
	}()

	app.Listen(PORT)
}
