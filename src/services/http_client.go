package services

import (
	"io"
	"net/http"
	"os"
)

func TeachableHttpGet(path string) ([]byte, error) {
	teachableBaseUrl := os.Getenv("TEACHABLE_BASE_URL")
	apiKey := os.Getenv("API_KEY")
	var error error
	client := &http.Client{}
	request, err := http.NewRequest("GET", teachableBaseUrl+path, nil)
	if err != nil {
		error = err
		return nil, error
	}
	request.Header = http.Header{
		"apiKey": {apiKey},
	}
	response, err := client.Do(request)
	if err != nil {
		error = err
		return nil, error
	}
	bytes, _ := io.ReadAll(response.Body)
	return bytes, nil
}
