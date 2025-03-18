package fiberredoc

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/jenggo/go-redoc"
)

func New(doc redoc.Redoc) fiber.Handler {
	return adaptor.HTTPHandlerFunc(doc.Handler())
}
