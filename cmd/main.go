package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Z33DD/Napoleon/db"
	"github.com/Z33DD/Napoleon/views"
	"github.com/go-redis/redis"
)

func main() {
	db.Client = getRedisClient()

	http.HandleFunc("/add", views.AddLink)
	http.HandleFunc("/s/", views.GetLink)

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	fmt.Println("Server lintening on http://" + host + ":" + port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}

func getRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	psswd := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))

	if err != nil {
		log.Fatal(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: psswd,
		DB:       db,
	})
	_, err = client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	return client
}
