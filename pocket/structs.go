package pocket

// Request request object for pocket get endpoint
type Request struct {
	ConsumerKey string `json:"consumer_key"`
	AccessToken string `json:"access_token"`
	State       string `json:"state"`
	ContentType string `json:"contentType"`
	Since       string `json:"since"`
}

// Response response object for pocket get endpoint
type Response struct {
	List map[string]Article `json:"list"`
}

// Pocket API
type Pocket struct {
	ConsumerKey  string
	AccessToken  string
	ReadingSpeed int // n words/minute
}

// Article Archived Pocket Article
type Article struct {
	TimeToRead    int    `json:"time_to_read"`
	ResolvedTitle string `json:"resolved_title"`
	ResolvedURL   string `json:"resolved_url"`
	WordCount     string `json:"word_count"`
}
