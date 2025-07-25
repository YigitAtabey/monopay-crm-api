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

	// Sadece admin tüm kullanıcıları görebilir, user sadece kendini, blocked erişemez.
	api.Get("/users", middleware.RequireAuth, controllers.GetUsers)

	api.Get("/profile", middleware.RequireAuth, controllers.Profile)

	// Görev işlemleri: blocked erişemez, user sadece kendi, admin hepsini yönetebilir
	api.Get("/tasks", middleware.RequireAuth, controllers.GetTasks)
	api.Post("/tasks", middleware.RequireAuth, controllers.CreateTask)
	api.Put("/tasks/:id", middleware.RequireAuth, controllers.UpdateTask)
	api.Delete("/tasks/:id", middleware.RequireAuth, controllers.DeleteTask)

	api.Post("/logout", middleware.RequireAuth, controllers.Logout)

	// Admin işlemleri: sadece admin yetkisiyle erişilebilir
	admin := app.Group("/api/admin", middleware.RequireAuth, middleware.RequireAdmin)
	admin.Put("/block/:id", controllers.BlockUser)        // Kullanıcıyı ID ile engelle
	admin.Put("/make", controllers.MakeAdminByEmail)      // Kullanıcıyı email ile admin yap
	admin.Put("/block", controllers.BlockUserByEmail)     // Kullanıcıyı email ile engelle
	admin.Put("/unblock", controllers.UnblockUserByEmail) // Kullanıcıyı email ile engeli kaldır <-- EKLENDİ!
}
