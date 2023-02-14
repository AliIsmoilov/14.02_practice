package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"app/models"
)

type userRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) GetAll(search string)([]models.User, error){

	var users []models.User

	b, err := ioutil.ReadFile("data/users.json")

	if err != nil{
		return nil, err
	}

	err = json.Unmarshal(b, &users)


	if err != nil{
		return nil, err
	}

	var found_users []models.User

	for _, user := range users {
		
		if strings.Contains(user.First_name, search)|| strings.Contains(user.Last_name, search){
			found_users = append(found_users, user)
		}
	}

	return found_users, nil
}

func (u *userRepo) Update(id string, UpdateUser models.User)error{

	var users []models.User

	b, err := ioutil.ReadFile("data/users.json")

	if err != nil{
		return err
	}

	err = json.Unmarshal(b, &users)


	if err != nil{
		return err
	}

	for in, user := range users{
		if user.Id == id{
			users[in] = UpdateUser

			body, err := json.MarshalIndent(users, "", "	")
			err = ioutil.WriteFile("data/users.json", body, os.ModePerm)

			if err != nil{
				return err
			}

			return nil
		}

	}

	return err
}


func (u *userRepo) GerById(id string) (found_user models.User, err error){

	var users []models.User

	b , err := ioutil.ReadFile("data/users.json")

	if err != nil{
		return found_user, err
	}

	err = json.Unmarshal(b, &users)

	if err != nil{
		return found_user, err
	}

	for _, user := range users{

		if user.Id == id{
			return user, nil
		}
	}

	return found_user, errors.New("Not found such id")

}


func (u *userRepo) Delete(id string) (error){

	var users []models.User

	b, err := ioutil.ReadFile("data/users.json")

	if err != nil{
		return err
	}

	err = json.Unmarshal(b, &users)

	if err != nil{
		return err
	}

	exist := false

	for in, user := range users{

		if user.Id == id{

			exist = true

			users = append(users[:in], users[in+1:]...)

			body, err := json.MarshalIndent(users, "", "	")
			err = ioutil.WriteFile("data/users.json", body, os.ModePerm)

			if err != nil{
				return err
			}

			return nil

		}
	}

	if !exist {
		return errors.New("Not found such user")
	}

	return nil
}
