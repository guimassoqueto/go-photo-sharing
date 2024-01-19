package main

import (
	"fmt"
	"github.com/google/uuid"
)

// GenerateUUID generates a random UUID and returns it as a string.
func GenerateUUIDV4() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("error generating UUID: %w", err)
	}
	return uuid.String(), nil
}

func main() {
	uuid, err := GenerateUUIDV4()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
	}
	fmt.Println(uuid)
}
