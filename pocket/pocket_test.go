package pocket_test

import (
	"testing"

	"github.com/akramsaouri/pocket-karma/pocket"
)

func TestMinRead(t *testing.T) {
	p := pocket.Pocket{ReadingSpeed: 125}
	articles := []pocket.Article{
		{TimeToRead: 10},
		{TimeToRead: 12},
		{WordCount: "500"},
		{TimeToRead: 8},
	}
	minRead, _ := p.MinRead(articles)
	if minRead != 34 {
		t.Errorf("MinRead was incorrect, got %d, want %d.", minRead, 34)
	}
}
