package helper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)

	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return namedQuery, args
}

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}
func ComparePasswords(hashedPass, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPass, pass)
}

var (
	counter int
	mutex   sync.Mutex
	rnd     *rand.Rand
)

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GenerateUniqueID generates a unique six-digit ID as a string
func GenerateUniqueID() string {
	mutex.Lock()
	defer mutex.Unlock()

	// Increment the counter
	counter++

	// Reset the counter if it exceeds 999,999
	if counter > 999999 {
		counter = 1
	}

	// Generate a random number between 0 and 9999 (inclusive)
	randomNumber := rnd.Intn(10000)

	// Combine the counter with the random number to form the ID
	id := fmt.Sprintf("%06d", counter*10000+randomNumber)

	return id
}
