package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/gofrs/uuid"
)

type CircleQueryModel struct {
	Search string    `query:"search"`
	Page   int64     `query:"page"`
	Owner  uuid.UUID `query:"owner"`
}

// QueryCircleHandle handle query on circle
func QueryCircleHandle(c *fiber.Ctx) error {
	fmt.Println("inside query circle handle")
	// Create service
	return nil
}

// GetMyCircleHandle handle get authed user circle
func GetMyCircleHandle(c *fiber.Ctx) error {

	// Create service
	fmt.Println("inside get my circle handle")
	return nil
}

// GetCircleHandle handle get a circle
func GetCircleHandle(c *fiber.Ctx) error {

	// Create service
	fmt.Println("inside get circle handle")
	return nil

}
