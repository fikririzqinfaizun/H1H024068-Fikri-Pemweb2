package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func WaktuProses(c fiber.Ctx) error {
	mulai := time.Now()

	err := c.Next()

	durasi := time.Since(mulai)

	log.Printf(
		"Request %s %s diproses selama %s",
		c.Method(),
		c.Path(),
		durasi,
	)

	return err
}
