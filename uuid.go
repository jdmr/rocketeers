package main

import (
	"crypto/rand"
	"fmt"
)

// UUID returns a string that can be used as a UUID
func UUID() (string, error) {
	n := 5
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("Could not generate UUID %v", err)
	}
	s := fmt.Sprintf("%X", b)
	return s, nil
}
