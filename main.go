package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/akramsaouri/pocket-karma/pocket"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type response struct {
	Success      bool             `json:"success"`
	Articles     []pocket.Article `json:"articles,omitempty"`
	TotalMinRead int              `json:"total_minutes_read,omitempty"`
	Error        interface{}      `json:"error,omitempty"`
}

func index(w http.ResponseWriter, r *http.Request) {
	readingSpeed, _ := strconv.Atoi(os.Getenv("READING_SPEED"))
	p := pocket.Pocket{
		ConsumerKey:  os.Getenv("POCKET_CONSUMER_KEY"),
		AccessToken:  os.Getenv("POCKET_ACCESS_TOKEN"),
		ReadingSpeed: readingSpeed,
	}
	articles, err := p.Articles("archive", "article", sinceFirstOfMonth())
	if err != nil {
		handleError(w, err)
		return
	}
	totalMinRead, err := p.MinRead(articles)
	if err != nil {
		handleError(w, err)
		return
	}
	resJSON(w, response{
		Success:      true,
		Articles:     articles,
		TotalMinRead: totalMinRead,
	})
}

func handleError(w http.ResponseWriter, err error) {
	resJSON(w, response{
		Success: false,
		Error:   err.Error(),
	})
}

func resJSON(w http.ResponseWriter, rs response) {
	j, e := json.Marshal(rs)
	if e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func sinceFirstOfMonth() string {
	now := time.Now()
	d := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	return strconv.FormatInt(d.Unix(), 10)
}
