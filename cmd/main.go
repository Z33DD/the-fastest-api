package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Z33DD/Napoleon/db"
	"github.com/Z33DD/Napoleon/views"
	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db.Client = getRedisClient()

	http.HandleFunc("/add", views.AddLink)
	http.HandleFunc("/s/", views.GetLink)

	port := os.Getenv("PORT")

	fmt.Println("Server lintening on http://0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func getRedisClient() *redis.Client {
	config, err := redis.ParseURL(os.Getenv("REDISTOGO_URL"))

	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(config)
	_, err = client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	return client
}
