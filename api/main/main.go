package main

import (
    "encoding/json"
    "net/http"
    "google.golang.org/appengine"
    
	bd "github.com/hamidzia/tracky/api/birthday" 
	log "google.golang.org/appengine/log"   
)

func init() {
	http.HandleFunc("/birthday", GetBirthdayHandler)
}

func main() {
    appengine.Main()
}

func GetBirthdayHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

    b, err := bd.ReadBirthdays(c, "data/birthdays.csv")
    if err != nil {
        log.Infof(c, "%v", err)
    }
    log.Infof(c, "Total Birthdays: %d", len(b))
    json.NewEncoder(w).Encode(b)    
}

