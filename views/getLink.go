package views

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Z33DD/Napoleon/db"
)

// GetLink - Find link that matches the shortened link in the linkList
func GetLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	shortUrl := pathArgs[2]
	original, err := db.Client.Get(shortUrl).Result()
	log.Printf("[Redirect] %s -> %s", shortUrl, original)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Does not exists!")
		return
	}
	log.Printf("Redirected from %s to: %s", shortUrl, original)
	http.Redirect(w, r, original, http.StatusPermanentRedirect)
	return
}
