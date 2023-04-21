package handlers

import (
	"fmt"
	"net/http"

	"example.com/core/pkg/log"
	"example.com/core/utils"
	"example.com/function/database"
	service "example.com/function/services"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/gofrs/uuid"
)

// DeleteCircleHandle handle delete a circle
func DeleteCircleHandle(c *fiber.Ctx) error {

	// params from /circles/:circleId

	circleId := c.Params("circleId")
	fmt.Println("circle id is : ", circleId)
	if circleId == "" {
		errorMessage := fmt.Sprintf("Circle Id is required!")
		log.Error(errorMessage)
		return c.Status(http.StatusBadRequest).JSON(utils.Error("circleIdRequired", errorMessage))
	}

	circleUUID, uuidErr := uuid.FromString(circleId)
	if uuidErr != nil {
		errorMessage := fmt.Sprintf("UUID Error %s", uuidErr.Error())
		log.Error(errorMessage)
		return c.Status(http.StatusBadRequest).JSON(utils.Error("circleIdIsNotValid", "Circle id is not valid!"))

	}
	fmt.Println("creating service")

	// Create service
	circleService, serviceErr := service.NewCircleService(database.Db)
	if serviceErr != nil {
		log.Error("NewCircleService %s", serviceErr.Error())
		return c.Status(http.StatusInternalServerError).JSON(utils.Error("internal/circleService", "Error happened while creating circleService!"))
	}

	// currentUser, ok := c.Locals(types.UserCtxName).(types.UserContext)
	// if !ok {
	// 	log.Error("[DeleteCircleHandle] Can not get current user")
	// 	return c.Status(http.StatusBadRequest).JSON(utils.Error("invalidCurrentUser",
	// 		"Can not get current user"))
	// }
	currentUser_UserID, uuidErr := uuid.FromString(circleId)
	if uuidErr != nil {
		errorMessage := fmt.Sprintf("UUID Error %s", uuidErr.Error())
		log.Error(errorMessage)
		return c.Status(http.StatusBadRequest).JSON(utils.Error("circleIdIsNotValid", "Circle id is not valid!"))

	}
	fmt.Println("calling delete circle by owner function")
	if err := circleService.DeleteCircleByOwner(currentUser_UserID, circleUUID); err != nil {
		errorMessage := fmt.Sprintf("Delete Circle Error %s - %s", circleUUID.String(), err.Error())
		log.Error(errorMessage)
		return c.Status(http.StatusBadRequest).JSON(utils.Error("deleteCircle", "Can not delete circle!"))
	}
	fmt.Println("cirlce is delted by owner")
	return c.SendStatus(http.StatusOK)
}
