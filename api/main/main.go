package main

import (
    "encoding/json"
    "log"
    "net/http"
    "google.golang.org/appengine"
    
    bd "github.com/hamidzia/rest-api/birthday"    
)

func init() {
	http.HandleFunc("/birthday", GetBirthdayHandler)
}

func main() {
    appengine.Main()
}

func GetBirthdayHandler(w http.ResponseWriter, r *http.Request) {
    b, err := bd.ReadBirthdays("../data/birthdays.csv")
    if err != nil {
        log.Printf("%v", err)
    }
    log.Printf("Total Birthdays: %d", len(b))
    json.NewEncoder(w).Encode(b)    
}

