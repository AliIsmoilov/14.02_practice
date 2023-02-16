package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type postRepo struct {
	fileName string
	file     *os.File
}


func NewPostRepo(fileName string, file *os.File) *postRepo {
	return &postRepo{
		fileName: fileName,
		file:     file,
	}
}

func (p *postRepo) CreatePost(Newpost models.CreatePost) (id string, err error){
	
	var posts []*models.Post
	err = json.NewDecoder(p.file).Decode(&posts)

	if err != nil{
		return "", err
	}

	id = uuid.New().String()

	posts = append(posts, &models.Post{
		Id: id,
		Title: Newpost.Title,
		Description: Newpost.Description,
		User: Newpost.UserId,
	})

	body, err := json.MarshalIndent(posts, "", "	")

	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)
	
	if err != nil{
		return "", err
	}

	return id, nil
}

func (p *postRepo) DeletePost(id string) (err error){

	var posts []*models.Post
	err = json.NewDecoder(p.file).Decode(&posts)

	if err != nil{
		return err
	}


	for ind, post := range posts{
		
		if post.Id == id{
		
			posts = append(posts[:ind], posts[ind+1:]...)

			body, err := json.MarshalIndent(posts, "", "	")
			err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

			if err != nil{
				return err
			}
			return nil

		}
	}

	return errors.New("Not found post with such Id")

}

func (p *postRepo) UpdatePost(UPost models.UpdatePost, id string) (err error){

	var posts []*models.Post
	err = json.NewDecoder(p.file).Decode(&posts)

	if err != nil{
		return err
	}

	for ind, post := range posts{

		if post.Id == id{

			posts[ind].Title = UPost.Title
			posts[ind].Description = UPost.Description

			body, err := json.MarshalIndent(posts, "", "	")
			err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

			if err != nil{
				return err
			}

			return nil
		}
	}

	return errors.New("Not found Post with such index")
}

func (p *postRepo) GetPostById(id string) (err error, Fpost models.Post){

	var posts []*models.Post
	err = json.NewDecoder(p.file).Decode(&posts)

	if err != nil{
		return err, Fpost
	}

	for _, post := range posts{
		if post.Id == id{
			
			Fpost = *post
		
			return nil, Fpost

		}
	}

	return errors.New("Not found Post with such ID"), Fpost

}


func (p *postRepo) GetAll() (posts []models.Post, err error){

	err = json.NewDecoder(p.file).Decode(&posts)

	if err != nil{
		return posts, err
	}

	return posts, nil
}