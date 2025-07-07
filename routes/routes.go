package routes

import (
	"monopay-crm-api/controllers"
	"monopay-crm-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)

	api.Get("/profile", middleware.RequireAuth, controllers.Profile)
	api.Get("/tasks", middleware.RequireAuth, controllers.GetTasks)

	// Yeni eklenen görev oluşturma route'u:
	api.Post("/tasks", middleware.RequireAuth, controllers.CreateTask)
	api.Put("/tasks/:id", middleware.RequireAuth, controllers.UpdateTask)
	api.Delete("/tasks/:id", middleware.RequireAuth, controllers.DeleteTask)

}
