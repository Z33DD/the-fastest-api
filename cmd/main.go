package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/Z33DD/Napoleon/db"
	"github.com/Z33DD/Napoleon/views"
	"github.com/go-redis/redis"
)

func main() {
	db.Client = getRedisClient()

	http.HandleFunc("/add", views.AddLink)
	http.HandleFunc("/s/", views.GetLink)

	port := os.Getenv("PORT")

	fmt.Println("Server lintening on http://0.0.0.0:" + port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

func getRedisClient() *redis.Client {
	config, err := getOptionsFromUrl(os.Getenv("REDISTOGO_URL"))

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

func getOptionsFromUrl(urlString string) (*redis.Options, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	password, _ := u.User.Password()

	options := &redis.Options{
		Addr:     u.Host,
		Password: password,
		DB:       0,
	}

	return options, nil
}
