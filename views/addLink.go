package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Z33DD/Napoleon/services"
)

func AddLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	key, ok := r.URL.Query()["link"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing parameters")
		return
	}
	original := key[0]

	shortUrl := services.GenShortenedLink(original)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	resp := make(map[string]string)
	resp["message"] = "Shortened link Created"
	resp["shortlink"] = shortUrl
	resp["original"] = original
	jsonResp, err := json.Marshal(resp)
	log.Printf("[Add] %s -> %s", shortUrl, original)

	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)

	return
}
