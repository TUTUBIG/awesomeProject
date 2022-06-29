package main

import (
	"fmt"
)
import "github.com/gofiber/fiber/v2"

func main() {
	svc := fiber.New(fiber.Config{
		Prefork:               false,
		ServerHeader:          "Ankr Power",
		DisableStartupMessage: true,
	})

	svc.Post("/test", test)

	svc.Post("/*", wichard)
	if err := svc.Listen(":8888"); err != nil {
		panic(err)
	}
}

func wichard(c *fiber.Ctx) error {
	fmt.Println("wichard")
	return nil
}

func test(c *fiber.Ctx) error {
	fmt.Println("test")
	return nil
}
