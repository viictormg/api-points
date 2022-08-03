package server

import (
	"api-points/models"

	"github.com/gofiber/fiber/v2"
)

func (s Server) RedeemPoints(ctx *fiber.Ctx) error {
	var req models.Redeem

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	err := s.Feature.RedeemPoints(req)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString("premio reclamado con exito ")

}
