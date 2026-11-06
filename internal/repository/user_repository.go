/*
 * Copyright (c) 2026 Itjen Kemenag. All rights reserved.
 * @author Yunus Bahtiar
 * Licensed under the Proprietary / Commercial License.
 */

package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/ybahtiarz/simtlhp-service/internal/domain"
)

type userRepository struct {
	db map[string]*domain.User
	mu sync.RWMutex
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		db: make(map[string]*domain.User),
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.db[user.ID] = user
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.db[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
