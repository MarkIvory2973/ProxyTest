package httpx

import (
	"encoding/json"
	"io"
	"net/http"
)

type HTTPResponse struct {
	Content string
}

func (response HTTPResponse) ToJson() (map[string]interface{}, error) {
	var jsonContent map[string]interface{}
	err := json.Unmarshal([]byte(response.Content), &jsonContent)

	return jsonContent, err
}

func HTTPGet(url string) (HTTPResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return HTTPResponse{}, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return HTTPResponse{}, err
	}

	return HTTPResponse{Content: string(content)}, nil
}
