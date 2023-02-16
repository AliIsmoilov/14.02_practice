package jsondb

import (
	"app/config"
	"app/storage"
	"os"
)

type Store struct {
	user *userRepo
	post *postRepo
}

func NewFileJson(cfg *config.Config)(storage.StorageI, error){

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	postFile, err := os.Open(cfg.Path+cfg.PostsFileName)

	if err != nil{
		return nil, err
	}

	return &Store{
		user: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		post: NewPostRepo(cfg.Path+cfg.PostsFileName, postFile),
	}, nil

}

func (s *Store) CloseDB() {
	s.user.file.Close()
}

func (s *Store) User() storage.UserRepoI {
	return s.user
}

func (s *Store) Post() storage.PostRepoI{
	return s.post
}