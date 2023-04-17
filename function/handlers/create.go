package handlers

import (
	"fmt"
	"net/http"

	"example.com/core/pkg/log"
	"example.com/core/types"
	"example.com/core/utils"
	"example.com/function/database"
	"example.com/function/models"
	"github.com/gofiber/fiber/v2"
	domain "github.com/red-gold/ts-serverless/micros/circles/dto"
	service "github.com/red-gold/ts-serverless/micros/circles/services"
)

const followingCircleName = "Following"

// CreateCircleHandle handle create a new circle
func CreateCircleHandle(c *fiber.Ctx) error {
	fmt.Println("inside create circle handle")
	// Create the model object
	model := new(models.CreateCircleModel)
	if err := c.BodyParser(model); err != nil {
		errorMessage := fmt.Sprintf("Parse CreateCircleModel Error %s", err.Error())
		log.Error(errorMessage)
		return c.Status(http.StatusInternalServerError).JSON(utils.Error("internal/parseModel", "Error happened while parsing model!"))
	}
	if model.Name == followingCircleName {
		errorMessage := fmt.Sprintf("Can not use 'Following' as a circle name")
		log.Error(errorMessage)
		return c.Status(http.StatusBadRequest).JSON(utils.Error("followingCircleNameIsReserved", errorMessage))
	}
	currentUser, ok := c.Locals(types.UserCtxName).(types.UserContext)
	if !ok {
		log.Error("[CreateCircleHandle] Can not get current user")
		return c.Status(http.StatusBadRequest).JSON(utils.Error("invalidCurrentUser",
			"Can not get current user"))
	}

	// Create a new circle
	newCircle := &domain.Circle{
		OwnerUserId: currentUser.UserID,
		Name:        model.Name,
		IsSystem:    false,
	}

	// Create service
	circleService, serviceErr := service.NewCircleService(database.Db)
	if serviceErr != nil {
		log.Error("NewCircleService %s", serviceErr.Error())
		return c.Status(http.StatusInternalServerError).JSON(utils.Error("internal/circleService", "Error happened while creating circleService!"))
	}

	if err := circleService.SaveCircle(newCircle); err != nil {
		errorMessage := fmt.Sprintf("Save Circle Error %s", err.Error())
		log.Error(errorMessage)
		return c.Status(http.StatusInternalServerError).JSON(utils.Error("internal/saveCircle", "Error happened while saving circle!"))
	}

	return c.JSON(fiber.Map{
		"objectId": newCircle.ObjectId.String(),
	})

	return nil
}

// CreateFollowingHandle handle create a new circle
func CreateFollowingHandle(c *fiber.Ctx) error {
	fmt.Println("inside create follow handle")
	// params from /circles/following/:userId
	return nil
}
