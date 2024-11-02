package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ImageService struct {
	accessKey string
}

func NewImageService() *ImageService {
	return &ImageService{
		accessKey: os.Getenv("UNSPLASH_ACCESS_KEY"),
	}
}

func (s *ImageService) GetRandomImage(query string) (string, error) {
	url := fmt.Sprintf("https://api.unsplash.com/photos/random?query=%s&client_id=%s", query, s.accessKey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	urls, ok := result["urls"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response format")
	}

	thumbnail, ok := urls["small"].(string)
	if !ok {
		return "", fmt.Errorf("thumbnail URL not found")
	}

	return thumbnail, nil
}
