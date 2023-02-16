package controller

import (
	"app/models"
	"fmt"
)

func (c *Controller) CreatePost(Newpost models.CreatePost) (id string, err error){

	id, err = c.store.Post().CreatePost(Newpost)

	if err != nil{
		return "", err
	}

	return id, nil

}

func (c *Controller) DeletePost(id string) (err error){

	err = c.store.Post().DeletePost(id)

	return err
}

func (c *Controller) UpdatePost(UPost models.UpdatePost, id string) (err error){

	err = c.store.Post().UpdatePost(UPost, id)

	return err
}

func (c *Controller) GetPostById(id string) (Outpost models.GetPostById, err error){

	err, post := c.store.Post().GetPostById(id)

	if err != nil{
		return Outpost, err
	}
	
	User, _ := c.store.User().GetById(post.User)
	fmt.Println(User)
	if err != nil{
		return Outpost, err
	}

	Outpost.Id = post.Id
	Outpost.Title = post.Title
	Outpost.Description = post.Description

	Outpost.User=User

	return Outpost, nil
}

func (c *Controller) GetAllPosts() (posts []models.Post, err error){

	posts, err = c.store.Post().GetAll()

	if err != nil{
		return posts, err
	}

	return posts, nil
}

func (c *Controller) GetUserPosts(id string) (Uposts models.GetUserPosts, err error){

	posts, err := c.store.Post().GetAll()

	for _, post := range posts{
		if post.User == id{
			Uposts.Posts = append(Uposts.Posts, post)
		}
	}

	if len(Uposts.Posts) > 0{
		Uposts.User, err = c.store.User().GetById(id)

		if err != nil{
			return Uposts, err
		}
	}

	
	return Uposts, nil
}