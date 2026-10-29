package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Environment Variables dari Infisical ===")

	for _, key := range strings.Split(os.Getenv("INFISICAL_KEYS"), ",") {
		key = strings.TrimSpace(key)
		if key == "" {
			continue
		}
		value, ok := os.LookupEnv(key)
		if !ok {
			fmt.Printf("Key: %-20s | TIDAK ADA\n", key)
			continue
		}
		fmt.Printf("Key: %-20s | Value: %s\n", key, value)
	}
}