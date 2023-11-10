package test

import (
	"app/models"
	"math/rand"
	"net/http"
	"time"

	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestCreateBranch(t *testing.T) {
	response := &models.Branch{}

	request := &models.CreateBranch{
		Name:      faker.FirstName(),
		Address:   faker.LastName(),
		FoundedAt: generateRandomYear(1900, 2023),
	}

	resp, err := makeRequest(http.MethodPost, "/branch", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

}

// Generate a random year between startYear and endYear (inclusive)
func generateRandomYear(startYear, endYear int) int {
	rand.Seed(time.Now().UnixNano())

	year := rand.Intn(endYear-startYear+1) + startYear
	return year
}
