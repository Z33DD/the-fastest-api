package services

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/Z33DD/Napoleon/db"
)

func GenShortenedLink(original string) string {
	linkId := fmt.Sprint(rand.Int63n(1000))
	host := os.Getenv("HOST")
	db.Client.Set(linkId, original, 0)
	return "https://" + host + "/s/" + linkId
}
