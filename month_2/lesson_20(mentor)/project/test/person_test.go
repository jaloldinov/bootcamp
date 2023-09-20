package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"playground/cpp-bootcamp/models"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	response := &models.Person{}

	request := &models.CreatePerson{
		Name: faker.FirstName(),
		Job:  faker.LastName(),
		Age:  rand.Intn(100),
	}

	resp, err := makeRequest(http.MethodPost, "/person", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)
}
