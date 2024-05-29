package main

import (
	"github.com/OhMinsSup/hashnode-llm/cmd"
)

func main() {
	// app := fiber.New()

	// // Define a route for the GET method on the root path '/'
	// app.Get("/", func(c fiber.Ctx) error {
	// 	// Send a string response to the client
	// 	return c.SendString("Hello, World 👋!")
	// })

	// log.Fatal(app.Listen(":3000"))
	cmd.Execute()
}
