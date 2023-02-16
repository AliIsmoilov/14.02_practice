package config



type Config struct {
	Path string

	UserFileName string
	PostsFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/users.json"
	cfg.Path = "./data"
	cfg.PostsFileName = "/posts.json"

	return cfg
}