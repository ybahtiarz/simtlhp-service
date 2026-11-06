/*
 * Copyright (c) 2026 Itjen Kemenag. All rights reserved.
 * @author Yunus Bahtiar
 * Licensed under the Proprietary / Commercial License.
 */

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ybahtiarz/simtlhp-service/internal/handler"
	"github.com/ybahtiarz/simtlhp-service/internal/repository"
	"github.com/ybahtiarz/simtlhp-service/internal/service"
)

func main() {
	app := fiber.New()

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api/v1")
	api.Post("/users", userHandler.CreateUser)
	api.Get("/users/:id", userHandler.GetUserByID)

	log.Println("Server running on port :3000")
	log.Fatal(app.Listen(":3000"))
}
