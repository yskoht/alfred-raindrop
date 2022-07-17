package raindrop

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GetRaindropsBody struct {
	Result bool       `json:"result"`
	Items  []Raindrop `json:"items"`
}

type Raindrop struct {
	ID    int    `json:"_id"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

const (
	ENDPOINT = "https://api.raindrop.io/rest/v1"
)

func getRaindropsEndpoint(
	collectionId int,
	page int,
	perpage int,
) string {
	return fmt.Sprintf(
		"%s/%s/%d?page=%d&perpage=%d",
		ENDPOINT,
		"/raindrops",
		collectionId,
		page,
		perpage,
	)
}

func getRaindrops(token string, page int) ([]Raindrop, error) {
	const (
		COLLECTION_ID = 0
		PERPAGE       = 50
	)

	url := getRaindropsEndpoint(COLLECTION_ID, page, PERPAGE)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", token)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	_body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var body GetRaindropsBody
	json.Unmarshal(_body, &body)
	if !body.Result {
		return nil, errors.New("failed to fetch bookmark")
	}

	return body.Items, nil
}

func GetAllRaindrops(token string) ([]Raindrop, error) {
	slice := make([]Raindrop, 0)

	for page := 0; ; page++ {
		raindrops, err := getRaindrops(token, page)
		if err != nil {
			return nil, err
		}
		if len(raindrops) == 0 {
			break
		}
		slice = append(slice, raindrops...)
	}

	return slice, nil
}
