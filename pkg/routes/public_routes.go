package routes

import (
	"github.com/SergeyRyskin/makromusic-simulation/app/models/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
    // Create routes group.
    route := a.Group("/api/v1")

    // Routes for GET method:
    route.Get("/users", controllers.GetUsers)              // get list of all users
    route.Get("/user/:id", controllers.GetUser)            // get one user by ID
    route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}