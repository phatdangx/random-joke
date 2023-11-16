package model

type RandomJoke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []string `json:"categories"`
		ID         int      `json:"id"`
		Joke       string   `json:"joke"`
	} `json:"value"`
}
