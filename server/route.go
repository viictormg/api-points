package server

import "github.com/gofiber/fiber/v2"

func (s Server) RegisterRoute(app *fiber.App) {
	routes := app.Group("/api")
	routes.Post("/points", s.SavePoint)
	routes.Get("/historyPoints", s.GetHistoryPoints)
	routes.Get("/awards", s.GetAwards)
	routes.Post("/redeem", s.RedeemPoints)

}
