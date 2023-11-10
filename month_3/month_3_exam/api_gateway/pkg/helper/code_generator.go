package helper

import (
	"encoding/json"
	"net/smtp"
	"strconv"
)

func GenerateCode() string {
	// code := 1000 + rand.Intn(9000)
	codeString := strconv.Itoa(1000)
	return codeString
}

func MarshalToStruct(data interface{}, resp interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, resp)
	if err != nil {
		return err
	}

	return nil
}

func SendMail(email string, message string) (string, error) {
	auth := smtp.PlainAuth(
		"",
		"jaloldinovuz@gmail.com",
		"jaloldinovs",
		"omamoh@gmail.com",
	)

	if err := smtp.SendMail(
		"omamoh.gmail.com:555",
		auth,
		"jaloldinovuz@gmail.com",
		[]string{email},
		[]byte(message),
	); err != nil {
		return "try again", err
	}

	return "message sent", nil
}
