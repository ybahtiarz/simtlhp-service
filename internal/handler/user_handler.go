/*
 * Copyright (c) 2026 Itjen Kemenag. All rights reserved.
 * @author Yunus Bahtiar
 * Licensed under the Proprietary / Commercial License.
 */

package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ybahtiarz/simtlhp-service/internal/domain"
)

type UserHandler struct {
	userService domain.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}

// POST /api/v1/users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req domain.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.userService.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user,
	})
}

// GET /api/v1/users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userService.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": user,
	})
}
