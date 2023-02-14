package storage

import(
	"os"
	"app/config"
)

type Store struct {
	User *userRepo
}

func NewFileJson(cfg *config.Config)(*Store, error){

	userFile, err := os.Open(cfg.Path + cfg.UserFileName)

	if err != nil{
		return nil, err
	}

	return &Store{
		User: NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
	}, nil

}