package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}
func ComparePasswords(hashedPass, pass []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPass, pass)
}
