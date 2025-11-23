package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func generatePassword(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func main() {
	length := flag.Int("length", 16, "password")
	flag.Parse()
	password, err := generatePassword(*length)
	if err != nil {
		fmt.Println("Error generating password:", err)
		os.Exit(1)
	}
	fmt.Println(password)
}

