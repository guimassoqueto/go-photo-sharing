package utils

import (
	"fmt"

	"github.com/google/uuid"
)

// UUIDV4 generates a random UUID version 4 and returns it as a string.
func UUIDV4() (string, error) {
    uuid, err := uuid.NewRandom()
    if err != nil {
        return "", fmt.Errorf("error generating UUID: %w", err)
    }
    return uuid.String(), nil
}