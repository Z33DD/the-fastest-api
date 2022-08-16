package views

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Z33DD/Napoleon/db"
	"github.com/Z33DD/Napoleon/services"
)

// GetLink - Find link that matches the shortened link in the linkList
func GetLink(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	linkId := pathArgs[2]
	original, err := db.Client.Get(linkId).Result()
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Does not exists!")
		return
	}
	go services.Track(linkId)
	log.Printf("[Redirect] %s -> %s", linkId, original)
	http.Redirect(w, r, original, http.StatusPermanentRedirect)
}
