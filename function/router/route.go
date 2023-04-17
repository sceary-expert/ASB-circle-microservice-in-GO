// Copyright (c) 2021 Amirhossein Movahedi (@qolzam)
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package router

import (
	"example.com/function/handlers"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {

	// Middleware
	// authHMACMiddleware := func(hmacWithCookie bool) func(*fiber.Ctx) error {
	// 	var Next func(c *fiber.Ctx) bool
	// 	if hmacWithCookie {
	// 		Next = func(c *fiber.Ctx) bool {
	// 			if c.Get(types.HeaderHMACAuthenticate) != "" {
	// 				return false
	// 			}
	// 			return true
	// 		}
	// 	}
	// 	return authhmac.New(authhmac.Config{
	// 		Next:          Next,
	// 		PayloadSecret: *config.AppConfig.PayloadSecret,
	// 	})
	// }

	// authCookieMiddleware := func(hmacWithCookie bool) func(*fiber.Ctx) error {
	// 	var Next func(c *fiber.Ctx) bool
	// 	if hmacWithCookie {
	// 		Next = func(c *fiber.Ctx) bool {
	// 			if c.Get(types.HeaderHMACAuthenticate) != "" {
	// 				return true
	// 			}
	// 			return false
	// 		}
	// 	}
	// 	return authcookie.New(authcookie.Config{
	// 		Next:         Next,
	// 		JWTSecretKey: []byte(*config.AppConfig.PublicKey),
	// 	})
	// }

	// hmacCookieHandlers := []func(*fiber.Ctx) error{authHMACMiddleware(false), authCookieMiddleware(false)}

	// Routers
	app.Post("/following/:userId", handlers.CreateFollowingHandle)
	app.Post("/create-circle", handlers.CreateCircleHandle)
	app.Put("/", handlers.UpdateCircleHandle)
	app.Delete("/:circleId", handlers.DeleteCircleHandle)
	app.Get("/my", handlers.GetMyCircleHandle)
	app.Get("/id/:circleId", handlers.GetCircleHandle)
}
