package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"avatar-generator/pkg/avatar"
)

func main() {
	// Define and parse command-line flags
	country := flag.String("country", "china", "The country for the avatar (e.g., 'china', 'usa').")
	gender := flag.String("gender", "male", "The gender for the avatar (e.g., 'male', 'female').")
	outputDir := flag.String("output", "output", "The directory to save the generated avatar.")
	flag.Parse()

	// Validate inputs
	if *country == "" || *gender == "" {
		log.Fatal("Country and gender flags must be provided.")
	}

	fmt.Printf("Generating SVG avatar for country: %s, gender: %s\n", *country, *gender)

	// Create the generator
	g := avatar.NewGenerator(*country, *gender)

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