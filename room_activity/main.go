package main

import (
	"room_activity/server"
)

func main() {
	server.NewServer().Serve()
}