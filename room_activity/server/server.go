package server

import (
	"room_activity/blueprint"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	fiber *fiber.App
	blueprint *blueprint.Blueprint
}

func NewServer() *Server {
	return &Server{
		fiber: fiber.New(),
		blueprint: new(blueprint.Blueprint),
	}
}

func (server *Server) Serve() error {
	server.fiber.Post("/blueprint", server.postBlueprint)
	server.fiber.Post("/rooms/:room/exit", server.postRoomExit)
	server.fiber.Post("/rooms/:room/enter", server.postRoomEnter)

	return server.fiber.Listen(":3000")
}