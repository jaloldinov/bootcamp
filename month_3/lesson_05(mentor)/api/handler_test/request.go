package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func makeRequest(method, path string, req, res interface{}) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), method, fmt.Sprintf("%s%s", "http://localhost:8000", path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, res)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
