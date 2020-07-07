package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)



func main() {
	password1 := "123"
	password2 := "123"
	hash ,_:= bcrypt.GenerateFromPassword([]byte(password1), 14)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password2))
	fmt.Println(err)
}
