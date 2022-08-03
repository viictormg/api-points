package server

import "github.com/gofiber/fiber/v2"

func (s Server) RegisterRoute(app *fiber.App) {
	routes := app.Group("/api")

	// routes.Post("/redeem", s.GetProduct)
	// routes.Get("/balance")
	routes.Post("/points", s.SavePoint)

	routes.Get("/awards", s.GetAwards)

	routes.Post("/redeem", s.RedeemPoints)

}
