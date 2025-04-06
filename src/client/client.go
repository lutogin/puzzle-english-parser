package peClient

import (
	"bytes"
	"encoding/json"
	"go-pe-parser/src/config"
	"io"
	"net/http"
)

type PuzzleEnglishClient struct {
	cookies string
	config  *config.Config
}

type Response struct {
	ListWords string `json:"listWords"`
}

func NewPuzzleEnglishClient(cookies string, cfg *config.Config) (*PuzzleEnglishClient, error) {
	return &PuzzleEnglishClient{
		cookies: cookies,
		config:  cfg,
	}, nil
}

func (c *PuzzleEnglishClient) GetDictionaryPage(page int) (string, error) {
	url := c.config.GetDictionaryEndpoint()
	queryParams := c.config.GetDictionaryQueryParams(page)

	req, err := http.NewRequest("POST", url, bytes.NewBufferString(queryParams))
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("cookie", c.cookies)
	req.Header.Set("origin", c.config.API.BaseURL)
	req.Header.Set("referer", c.config.GetDictionaryEndpoint())
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jsonResp Response
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return "", err
	}

	return jsonResp.ListWords, nil
}
