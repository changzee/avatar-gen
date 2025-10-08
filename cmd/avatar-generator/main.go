package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"strings"

	"avatar-generator/pkg/avatar"
	"avatar-generator/pkg/countries"
)

func main() {
	// Define and parse command-line flags
	country := flag.String("country", "cn", "The ISO 3166-1 alpha-2 country code for the avatar (e.g., 'cn', 'us').")
	gender := flag.String("gender", "male", "The gender for the avatar (e.g., 'male', 'female').")
	outputDir := flag.String("output", "output", "The directory to save the generated avatar.")
	flag.Parse()

	// Validate inputs
	processedCountry := strings.ToLower(*country)
	if !countries.IsValidCode(processedCountry) {
		log.Fatalf("Invalid country code: %s. Please provide a valid ISO 3166-1 alpha-2 code.", *country)
	}

	if *gender == "" {
		log.Fatal("Gender flag must be provided.")
	}

	fmt.Printf("Generating SVG avatar for country: %s, gender: %s\n", processedCountry, *gender)

	// Create the generator
	g := avatar.NewGenerator(processedCountry, *gender)

	// Generate the avatar SVG as a string
	svgContent, err := g.Generate()
	if err != nil {
		log.Fatalf("Failed to generate avatar: %v", err)
	}

	// Ensure the output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Create a unique filename with the .svg extension
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("avatar_%s_%s_%d.svg", *country, *gender, timestamp)
	filePath := filepath.Join(*outputDir, fileName)

	// Write the SVG string content to the file
	err = ioutil.WriteFile(filePath, []byte(svgContent), 0644)
	if err != nil {
		log.Fatalf("Failed to save avatar to file: %v", err)
	}

	fmt.Printf("Successfully generated avatar and saved it to %s\n", filePath)
}