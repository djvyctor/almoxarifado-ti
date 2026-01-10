package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "admin123"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Hash gerado:")
	fmt.Println(string(hash))
}

// HASH GERADO:
// $2a$10$5O/cvhiz.fLdjUraCDvYaOlSlL8g7fkm5C.SRcfn5HdeF4DL2r7vC
