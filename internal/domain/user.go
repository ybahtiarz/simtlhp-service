/*
 * Copyright (c) 2026 Itjen Kemenag. All rights reserved.
 * @author Yunus Bahtiar
 * Licensed under the Proprietary / Commercial License.
 */

package domain

import "context"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=50"`
	Email string `json:"email" validate:"required,email"`
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id string) (*User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*User, error)
	GetUserByID(ctx context.Context, id string) (*User, error)
}
