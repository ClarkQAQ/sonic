package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password, err := bcrypt.GenerateFromPassword([]byte("clarkqaq"), bcrypt.DefaultCost)
	if err != nil {
		panic("encrypt password" + err.Error())
	}
	fmt.Println(string(password))
}
