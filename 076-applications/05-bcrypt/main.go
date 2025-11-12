package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	p := "password123"

	bp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Println("bcrypted password:", string(bp))

	loginPassword := "password123"
	err = bcrypt.CompareHashAndPassword(bp, []byte(loginPassword))
	if err != nil {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("Login successful")
	}
}