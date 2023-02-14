package main

import (
	"fmt"
	"log"

	"app/config"
	"app/controller"
	"app/storage"
)


func main(){

	// adress := models.Adress{
	// 	Street: "5911 Vella Isle",
    //     City: "Tallinn",
	// }

	// user := models.User{
	// 	Id: "c57aa672-902f-44c8-af9d-dfa02f62541a",
	// 	First_name: "Ali",
	// 	Last_name: "Ismoilov",
	// 	Gender: "Male",
	// 	Adress: adress,
	// }

	cfg := config.Load()

	store, err := storage.NewFileJson(&cfg)

	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, store)


	// found_user, err := c.GetById("c57aa672-902f-44c8-af9d-dfa02f62541a")

	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(found_user)

	
	err = c.Delete("32a803b5-3f1c-495e-8ce7-4649e4cbe3b1")

	if err != nil{
		log.Println(err)
	} else {
		fmt.Println("User has been deleted")
	}

}
