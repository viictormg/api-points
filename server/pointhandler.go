package server

import (
	"api-points/models"

	"github.com/gofiber/fiber/v2"
)

func (s Server) SavePoint(ctx *fiber.Ctx) error {

	var req models.Point

	req.IDShop = "d2683200-12d0-11ed-ab31-9828a64b41f5"

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	err := s.Feature.SavePoint(req)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).SendString("puntos cargados con exito ")
}

func (s Server) GetHistoryPoints(ctx *fiber.Ctx) error {
	idCustomer := "1d1326a8-12d5-11ed-ab31-9828a64b41f5"
	response, err := s.Feature.GetHistoryPoints(idCustomer)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
