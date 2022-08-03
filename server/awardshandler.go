package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s Server) GetAwards(ctx *fiber.Ctx) error {

	awards, err := s.Feature.GetAwards()

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(awards)

}
