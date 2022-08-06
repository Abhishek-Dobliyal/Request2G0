package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type request struct {
	Url         string      `json:"url"`
	HTTPMethod  string      `json:"httpMethod"`
	Json        interface{} `json:"json"`
	CustomParam interface{} `json:"customParams"`
}

type jsonTextResponse struct {
	Message    string
	StatusCode int
}

func fetchResponse(c *fiber.Ctx) error {
	req := new(request)
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error:", err.Error())
		return c.JSON(jsonTextResponse{Message: "error", StatusCode: 404})
	}
	
	return c.JSON(jsonTextResponse{Message: "success", StatusCode: 200})
}

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Post("/fetch-response", fetchResponse)
	log.Fatal(app.Listen(":8000"))
}
