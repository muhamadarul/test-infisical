package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Daftar Environment Variables dari Infisical ===")
	for _, env := range os.Environ() {
		// Filter agar hanya menampilkan key spesifik/semua key
		pair := strings.SplitN(env, "=", 2)
		fmt.Printf("Key: %-20s | Value: %s\n", pair[0], pair[1])
	}
}