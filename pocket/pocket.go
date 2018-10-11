package pocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

const endpoint = "https://getpocket.com/v3/get"
const headers = "application/json; charset=utf-8"

// Articles calls pocket and fetch articles
func (p Pocket) Articles(state, contentType, since string) ([]Article, error) {
	bs := new(bytes.Buffer)
	json.NewEncoder(bs).Encode(Request{
		AccessToken: p.AccessToken,
		ConsumerKey: p.ConsumerKey,
		State:       state,
		ContentType: contentType,
		Since:       since,
	})
	res, _ := http.Post(endpoint, headers, bs)
	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		// Return early with error body on failure
		return nil, errors.New(string(b))
	}
	r := Response{}
	err := json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	articles := []Article{}
	// convert map[string]Article to []Article
	for _, v := range r.List {
		articles = append(articles, v)
	}
	return articles, nil
}

// MinRead computes and returns totla minuntes read on articles
func (p Pocket) MinRead(articles []Article) (int, error) {
	minRead := 0
	for _, article := range articles {
		// compute minutes read
		if article.TimeToRead > 0 {
			// work with time_to_read if available
			minRead += article.TimeToRead
		} else {
			// manually compute reading minutes using word_count and reading speed
			wc, err := strconv.Atoi(article.WordCount)
			if err != nil {
				return 0, err
			}
			minRead += wc / p.ReadingSpeed
		}
	}
	return minRead, nil
}
