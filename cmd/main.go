package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}

	c := controller.NewController(&cfg, jsondb)

	// c.CreateUser(&models.CreateUser{Name: "Ali", Surname: "Ismoilov", Balance: 500000})
	// p, err := c.AddShopCart(&models.AddShopCart{
	// 	ProductId: "48b934e9-ed15-4779-8d0d-e45c61c7a089",
	// 	UserId: "c6772cfd-f356-499d-a03b-75e76630b719",
	// 	Count: 6,
	// 	},
	// )

	// if err != nil{
	// 	log.Println(err)
	// 	return
	// }

	// fmt.Println(p)

	
	// userID := "c6772cfd-f356-499d-a03b-75e76630b719"


	// total, e := c.CalcTotalPrice(models.CalculateShop{
	// 	UserID:         userID,
	// 	Discount:       0,
	// 	DiscountStatus: "precent",
	// })
	// if e != nil {
	// 	log.Fatal(e)
	// }

	// fmt.Println("Total price:", total)

	// err = c.WithdrawUserBalance(userID, total)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = c.SendMoney(&models.ExchangeMoney{
		SenderID: "c6772cfd-f356-499d-a03b-75e76630b719",
		RecieverID: "5b899928-e98e-4597-94b3-78bc98272533",
		Summa: 10000,
	})

	if err != nil{
		log.Println(err)
		return
	}
	fmt.Println("Money sended successfully")

	// err = c.CreateKomissiya(models.Komissiya{Balance: 0})
}