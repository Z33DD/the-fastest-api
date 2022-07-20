package services

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Z33DD/Napoleon/db"
)

func GenShortenedLink(original string) string {
	shortened := fmt.Sprint(rand.Int63n(1000))
	host := os.Getenv("HOST")
	db.Client.Set(shortened, original, 0)
	return "https://" + host + "/s/" + shortened
}
