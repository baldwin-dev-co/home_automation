package server

import "github.com/gofiber/fiber/v2"

func (server *Server) postRoomEnter(c *fiber.Ctx) error {
	room, exists := (*server.blueprint)[c.Params("room")]
	if !exists {
		return c.SendStatus(404)
	}

	from := c.Query("from")
	if from == "" && !room.Entrance {
		return c.Status(400).SendString("Non entrances must include a `from` query parameter")
	}

	err := room.Enter(from)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(200)
}