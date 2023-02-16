package models

type User struct{
	Id string `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Gender string `json:"gender"`
	Adress Adress `json:"adress"`
	Friends []Friend `json:"friends"`
}

type Adress struct{
	Street string `json:"street"`
	City string `json:"city"`
}

type Friend struct{
	Id string `json:"id"`
	Email string `json:"email"`
	Phone_number string `json:"phone_number"`
}

type UpdateUser struct {
	Id string `json:"id"`
	First_name string `json:"first_name"`
	Last_name string `json:"last_name"`
	Gender string `json:"gender"`
	Adress Adress `json:"adress"`
	Friends []Friend `json:"friends"`
}