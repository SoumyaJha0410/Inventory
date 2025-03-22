package main

import (
	"context"
	"log"
	"os"

	"github.com/SoumyaJha0410/backend/pkg/common/app"
	"github.com/SoumyaJha0410/backend/pkg/common/postgresql"
	"github.com/SoumyaJha0410/backend/pkg/controller"
	"github.com/SoumyaJha0410/backend/pkg/repository"
	"github.com/SoumyaJha0410/backend/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ctx := context.Background()
	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgresqlConfig)

	userRepository := repository.NewUserRepository(dbPool)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	productRepository:= repository.NewProductRepository(dbPool)
	productService:= service.NewProductService(productRepository)
	productController:=controller.NewProductController(productService)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{echo.HeaderContentType, echo.HeaderAuthorization},
		AllowCredentials: true,
	}))

	userController.RegisteredUserRoutes(e)
	productController.RegisterProductRoutes(e)
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on port", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
