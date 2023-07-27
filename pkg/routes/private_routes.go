package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/SergeyRyskin/makromusic-simulation/app/controllers"
	"github.com/SergeyRyskin/makromusic-simulation/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
    // Create routes group.
    route := a.Group("/api/v1")

    // Routes for POST method:
    route.Post("/user", middleware.JWTProtected(), controllers.CreateUser) // create a new user

    // Routes for PUT method:
    route.Put("/user", middleware.JWTProtected(), controllers.UpdateUser) // update one user by ID

    // Routes for DELETE method:
    route.Delete("/user", middleware.JWTProtected(), controllers.DeleteUser) // delete one user by ID
}