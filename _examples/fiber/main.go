package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/jenggo/go-redoc"
	fiberredoc "github.com/jenggo/go-redoc/fiber"
)

func main() {
	doc := redoc.Redoc{
		Title:       "Example API",
		Description: "Example API Description",
		SpecFile:    "./openapi.json",
		SpecPath:    "/openapi.json",
		DocsPath:    "/docs",
	}

	r := fiber.New()
	r.Use(fiberredoc.New(doc))

	println("Documentation served at http://127.0.0.1:8000/docs")
	panic(r.Listen(":8000"))
}
