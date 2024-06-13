package routes

import (
	"karyawan/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Grouping employee routes
	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", handlers.CreateEmployee)
	employeeGroup.Put("/:id", handlers.UpdateEmployee)
	employeeGroup.Delete("/:id", handlers.DeleteEmployee)
	employeeGroup.Get("/", handlers.GetAllEmployees)
	employeeGroup.Get("/:id", handlers.GetEmployeeByID)

	// Grouping position routes
	positionGroup := app.Group("/positions")
	positionGroup.Post("/", handlers.CreatePosition)
	positionGroup.Put("/:id", handlers.UpdatePosition)
	positionGroup.Delete("/:id", handlers.DeletePosition)
	positionGroup.Get("/", handlers.GetAllPositions)
	positionGroup.Get("/:id", handlers.GetPositionByID)

	// Grouping department routes
	departmentGroup := app.Group("/departments")
	departmentGroup.Post("/", handlers.CreateDepartment)
	departmentGroup.Put("/:id", handlers.UpdateDepartment)
	departmentGroup.Delete("/:id", handlers.DeleteDepartment)
	departmentGroup.Get("/", handlers.GetAllDepartments)
	departmentGroup.Get("/:id", handlers.GetDepartmentByID)
}
