package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var estoniaLoc *time.Location

func init() {
	var err error
	estoniaLoc, err = time.LoadLocation("Europe/Tallinn")
	if err != nil {
		log.Fatalf("could not load Europe/Tallinn timezone: %v", err)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().In(estoniaLoc)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, now.Format("15:04:05"))
}

func main() {
	http.HandleFunc("/", timeHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
