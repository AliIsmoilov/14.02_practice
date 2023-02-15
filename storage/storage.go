package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (int, error)
	GetPkey(*models.UserPrimaryKey) (*models.User, error)
}
