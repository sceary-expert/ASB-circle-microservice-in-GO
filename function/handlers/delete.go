package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// DeleteCircleHandle handle delete a circle
func DeleteCircleHandle(c *fiber.Ctx) error {
	fmt.Println("inside delete circle handle")
	return nil
}
