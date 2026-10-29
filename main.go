package main

import (
	"fmt"
	"os"
)

func main() {
	// Membaca variabel environment yang disuntikkan oleh Infisical
	dbHost := os.Getenv("DB_HOST")
	apiKey := os.Getenv("API_KEY")

	fmt.Println("=== Aplikasi Go Berhasil Dijalankan ===")
	fmt.Println("DB_HOST dari Infisical :", dbHost)
	fmt.Println("API_KEY dari Infisical :", apiKey)
}