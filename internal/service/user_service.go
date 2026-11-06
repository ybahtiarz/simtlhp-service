/*
 * Copyright (c) 2026 Itjen Kemenag. All rights reserved.
 * @author Yunus Bahtiar
 * Licensed under the Proprietary / Commercial License.
 */

package service

import (
	"context"
	"fmt"
	"time"

	"github.com/ybahtiarz/simtlhp-service/internal/domain"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{userRepo: repo}
}

func (s *userService) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	newUser := &domain.User{
		ID:    fmt.Sprintf("usr-%d", time.Now().UnixNano()),
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.userRepo.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}
