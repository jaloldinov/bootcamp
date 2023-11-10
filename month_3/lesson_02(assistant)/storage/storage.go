package storage

import (
	"auth/models"
	"context"
)

type StorageI interface {
	User() UsersI
}

type UsersI interface {
	CreateUser(context.Context, *models.CreateUser) (string, error)
	GetUser(context.Context, *models.IdRequest) (*models.User, error)
	GetAllUser(context.Context, *models.GetAllUserRequest) (*models.GetAllUser, error)
	UpdateUser(context.Context, *models.User) (string, error)
	DeleteUser(context.Context, *models.IdRequest) (string, error)

	GetByLogin(context.Context, *models.LoginRequest) (*models.LoginDataRespond, error)
}
