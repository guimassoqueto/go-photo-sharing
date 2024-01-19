package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// for i, arg := range os.Args {
	// 	fmt.Printf("Arg %d = %s\n", i, arg)
	// }
	switch os.Args[1] {
	case "hash":
		hash(os.Args[2])
	case "compare":
		compare(os.Args[2], os.Args[3])
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}

// hash hashes a password
func hash(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return "", err
	}
	fmt.Printf("%s\n", string(bcryptPassword))
	return fmt.Sprintf("%s\n", string(bcryptPassword)), nil
}

// compare compares a password with a hash
func compare(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("Invalid password")
		return false
	}
	fmt.Println("Valid password")
	return true
}
