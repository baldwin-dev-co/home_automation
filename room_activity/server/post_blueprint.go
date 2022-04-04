package server

import (
	"room_activity/blueprint"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) postBlueprint(c *fiber.Ctx) error {
	bp := blueprint.MakeBlueprint()
	err := c.BodyParser(&bp)
	if  err != nil {
		return c.SendStatus(400)
	}

	server.blueprint = &bp
	return c.SendStatus(200)
}