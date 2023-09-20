package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/test-go/testify/assert"
)

// unit test
func TestSum(t *testing.T) {
	a, b := 4, 9

	res := sum(a, b)

	assert.Equal(t, 13, res)

}
func sum(a, b int) int {
	return a + b
}

type argsAndExpected struct {
	a, b, expected int
}

// table driven test
func TestSumTD(t *testing.T) {
	cases := []argsAndExpected{
		{a: 6, b: 3, expected: 9},
		{a: 0, b: 8, expected: 8},
		{a: -4, b: -7, expected: -10},
		{a: -4, b: 0, expected: -4},
	}
	for _, c := range cases {
		res := sum(c.a, c.b)
		assert.Equal(t, c.expected, res, "result is not equal to expected")
	}

}

//

//

//

//
//
//
//
//
//
//
//
//
//

func makeRequest(method, path string, req, res interface{}) (*http.Response, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	request, err := http.NewRequestWithContext(context.Background(), method, fmt.Sprintf("%s%s", "http://localhost:8080", path), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	request.Header.Add("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(resp_body, res)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
