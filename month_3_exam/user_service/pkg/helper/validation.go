package helper

import (
	"errors"
	"regexp"
	"strings"
)

/*
The password must be at least 8 characters long.
The password must contain at least one uppercase letter.
The password must contain at least one lowercase letter.
The password must contain at least one digit.
The password must contain at least one special character->

	<-(e.g., ~!@#$%^&*()_+{}|:"<>?/[];',.).
*/
func IsValidPassword(password string) bool {
	// Check if the password is at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Check if the password contains at least one uppercase letter
	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return false
	}

	// Check if the password contains at least one lowercase letter
	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		return false
	}

	// Check if the password contains at least one digit
	if !strings.ContainsAny(password, "0123456789") {
		return false
	}

	// Check if the password contains at least one special character
	// (non-alphanumeric character)
	match, _ := regexp.MatchString(`[~!@#$%^&*()_+{}|:"<>?/\[\]\;',\.]`, password)
	if !match {
		return match
	}

	return true
}

// IsValidPhone ...
func IsValidPhone(phone string) bool {
	r := regexp.MustCompile(`^\+998[0-9]{2}[0-9]{7}$`)
	return r.MatchString(phone)
}

// IsValidLogin ...
func IsValidLogin(login string) bool {
	r := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]{5,29}$`)
	return r.MatchString(login)
}

func ValidPinfl(pinfl string) error {
	if pinfl == "" {
		return errors.New("error application passport_pinfl requirement body to model")
	}
	pattern := regexp.MustCompile(`^([0-9]{14})$`)

	if !(pattern.MatchString(pinfl)) {
		return errors.New("passport_pinfl must be 14 digits")
	}
	return nil
}

func ValidPassportNumber(number string) error {
	if number == "" {
		return errors.New("error application passport_number requirement body to model")
	}
	pattern := regexp.MustCompile(`^([0-9]{7})$`)

	if !(pattern.MatchString(number)) {
		return errors.New("passport_number must be 7 digits")
	}
	return nil
}

// IsValidEmail ...
func IsValidEmail(email string) bool {
	r := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return r.MatchString(email)
}

// IsValidUUID ...
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func IsValidUUIDV1(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// IsValidPrice ...
func IsValidPrice(price string) bool {
	r := regexp.MustCompile(`^\d+$`)
	return r.MatchString(price)
}
