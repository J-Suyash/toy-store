package config

import "os"

type Config struct {
	MongoURI          string
	UnsplashAccessKey string
}

func New() *Config {
	return &Config{
		MongoURI:          os.Getenv("MONGODB_URI"),
		UnsplashAccessKey: os.Getenv("UNSPLASH_ACCESS_KEY"),
	}
}
