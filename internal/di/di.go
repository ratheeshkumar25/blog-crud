package di

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ratheeshkumar25/blog-crud-api/config"
	_ "github.com/ratheeshkumar25/blog-crud-api/docs"
	"github.com/ratheeshkumar25/blog-crud-api/internal/db"
	"github.com/ratheeshkumar25/blog-crud-api/internal/handlers"
	"github.com/ratheeshkumar25/blog-crud-api/internal/repositories"
	"github.com/ratheeshkumar25/blog-crud-api/internal/services"
	"github.com/ratheeshkumar25/blog-crud-api/utility"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Blog CRUD API
// @version 1.0
// @description This is a simple CRUD API for blog posts using Fiber.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

func Init() {
	cnf := config.LoadConfig()
	app := fiber.New()

	log := utility.InitLogger()
	log.Println("Starting Blog API Server...")
	dbConn := db.ConnectDB(cnf)
	log.Println("Database connected successfully.")
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("Request: %s %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	blogRepo := repositories.NewBlogRepository(dbConn)
	blogSvc := services.NewBlogServices(blogRepo)
	handlers.NewBlogHandler(app.Group("/api/v1"), blogSvc)

	log.Fatal(app.Listen(":8080"))
}
