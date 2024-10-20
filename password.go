package main

import (
	"math/rand/v2"
)

const (
	letters           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters = "!@$%&*()_+-=[]{},.:;<>/?~"
	numbers           = "0123456789"
)

func generatePassword(length int) string {

	availableCharacters := [3]string{letters, specialCharacters, numbers}
	password := make([]byte, length)

	for i := range password {
		dataSet := availableCharacters[rand.IntN(2)]
		password[i] = dataSet[rand.IntN(len(dataSet))]
	}

	return string(password)
}
