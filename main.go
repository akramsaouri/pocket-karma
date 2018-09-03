package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// PocketReqData Pocket API Request Object
type PocketReqData struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	State       string `json:"state"`
	ContentType string `json:"contentType"`
	Since       string `json:"since"`
}

// PocketResData Pocket API Response Object
type PocketResData struct {
	List map[string]PocketItem `json:"list"`
}

// PocketItem Pocket article
type PocketItem struct {
	WordCount     string `json:"word_count"`
	TimeToRead    int    `json:"time_to_read"`
	ResolvedTitle string `json:"resolved_title"`
}

// Report response Object
type Report struct {
	Status     string      `json:"status"`
	KP         int         `json:"karma_points"`
	ReportDate string      `json:"report_date"`
	Error      interface{} `json:"error"`
}

const wordsPerOneMin = 125
const karmaCoefficient = 10

func index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	d := time.Date(now.Year(), time.August, 1, 0, 0, 0, 0, time.UTC)
	since := strconv.FormatInt(d.Unix(), 10)
	pocketReqData := PocketReqData{
		ConsumerKey: os.Getenv("POCKET_CONSUMER_KEY"),
		AccessToken: os.Getenv("POCKET_ACCESS_TOKEN"),
		State:       "archive",
		ContentType: "article",
		Since:       since,
	}
	bs := new(bytes.Buffer)
	json.NewEncoder(bs).Encode(pocketReqData)
	res, _ := http.Post("https://getpocket.com/v3/get", "application/json; charset=utf-8", bs)
	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		rp := Report{
			Error: string(b),
		}
		resJSON(w, rp)
	} else {
		pocketResData := PocketResData{}
		json.NewDecoder(res.Body).Decode(&pocketResData)
		var readMin int
		for _, item := range pocketResData.List {
			if item.TimeToRead > 0 {
				readMin += item.TimeToRead
			} else {
				wc, err := strconv.Atoi(item.WordCount)
				if err != nil {
					log.Fatal(err)
				}
				readMin += wc / wordsPerOneMin
			}
		}
		kp := (readMin * karmaCoefficient) / 60
		rd := fmt.Sprintf("%s %d", now.Month(), now.Year())
		rp := Report{
			"success",
			kp,
			rd,
			nil,
		}
		resJSON(w, rp)
	}
}

func resJSON(w http.ResponseWriter, rp Report) {
	j, e := json.Marshal(rp)
	if e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
