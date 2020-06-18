package utils

import (
	"log"

	"github.com/gofiber/fiber"
)

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(ctx *fiber.Ctx, body interface{}) {
	if err := ctx.BodyParser(body); err != nil {
		log.Fatal(err)
	}
}
