package server

import (
	"github.com/gofiber/fiber/v2"
)

func (server *Server) postRoomExit(c *fiber.Ctx) error {
	room, exists := (*server.blueprint)[c.Params("room")]
	if !exists {
		return c.SendStatus(404)
	}

	to := c.Query("to")
	if to == "" && !room.Entrance {
		return c.Status(400).SendString("Non entrances must include a `to` query parameter")
	}

	err := room.Exit(to)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}