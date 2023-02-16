package models

type PrimaryKey struct{
	Id string `json:"id"`
}

type Post struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	User string
}
type GetPostById struct{
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	User User
}

type GetUserPosts struct{
	Posts []Post
	User User
}

type CreatePost struct{
	Title string `json:"title"`
	Description string `json:"description"`
	UserId string
}

type UpdatePost struct{
	Title string `json:"title"`
	Description string `json:"description"`
}

