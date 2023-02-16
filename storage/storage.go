package storage

import(
	"app/models"
)

type StorageI interface{
	CloseDB()
	User() UserRepoI
	Post() PostRepoI
}

type UserRepoI interface{
	GetAll(search string)([]models.User, error)
	Update(id string, UpdateUser models.User)error
	GetById(id string) (found_user models.User, err error)
	Delete(id string) (error)
}

type PostRepoI interface{
	CreatePost(Newpost models.CreatePost) (id string, err error)
	DeletePost(id string) (err error)
	UpdatePost(UPost models.UpdatePost, id string) (err error)
	GetPostById(id string) (err error, post models.Post)
	GetAll() (posts []models.Post, err error)
}