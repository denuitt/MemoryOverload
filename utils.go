package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

func GenerateRandomHex() (string, error) {
	randomBytes := make([]byte, 256)

	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	var randomHex string = hex.EncodeToString(randomBytes)

	return randomHex, nil
}
