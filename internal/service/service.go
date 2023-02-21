package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	requestURL      string = "https://api.chucknorris.io/jokes/random"
	NumberOfRequest int    = 20
)

type Message struct {
	IconURL string `json:"icon_url"`
	ID      string `json:"id"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

func GetResults(ctx context.Context) ([]Message, error) {
	results := make(map[string]Message)

	for len(results) < 25 {
		msg, err := getMessage(ctx)
		if err != nil {
			return nil, err
		}

		results[msg.ID] = *msg
	}

	rMessages := make([]Message, 0)
	for _, value := range results {
		rMessages = append(rMessages, value)
	}

	return rMessages, nil
}

func getMessage(ctx context.Context) (*Message, error) {
	client := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data := make([]byte, 0)
	if data, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var msg Message
	if err = json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}

	return &msg, nil
}
