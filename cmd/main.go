package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/storage/jsondb"
)


func main(){


	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)

	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, jsondb)


	
	posts, err :=  c.GetUserPosts("c46d8fe5-6eb4-4e5c-a09f-d523bc84d188")

	if err != nil{
		log.Println(err)
		return
	}
	
	fmt.Println(posts.Posts)
	fmt.Println(posts.User)
}
