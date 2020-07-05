package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetHashingCost(hashedPassword []byte) int {
	cost, _ := bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
	return cost
}

func main() {
	password1 := "sec1222222222222222222222222222222222222"
	password2 := "sec12"

	hash1 ,_ :=HashPassword(password1)
	hash2 ,_ := HashPassword(password2)
	fmt.Println(hash1)
	fmt.Println(hash2)
	fmt.Println(CheckPasswordHash(password1,hash2))
}
