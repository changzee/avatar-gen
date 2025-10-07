package avatar

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const (
	// AssetRoot is the root directory for all assets.
	AssetRoot = "assets"
	// CommonDir is a fallback directory for assets that are not country or gender specific.
	CommonDir = "common"
)

// Generator holds the configuration for creating an avatar.
type Generator struct {
	Country string
	Gender  string
}

// NewGenerator creates a new avatar generator with the specified options.
func NewGenerator(country, gender string) *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{
		Country: country,
		Gender:  gender,
	}
}

// Generate creates a new avatar and returns it as an SVG string.
func (g *Generator) Generate() (string, error) {
	layers, err := g.collectLayers()
	if err != nil {
		return "", fmt.Errorf("failed to collect layers: %w", err)
	}

	if len(layers) == 0 {
		return "", fmt.Errorf("no layers found for the given criteria (country: %s, gender: %s)", g.Country, g.Gender)
	}

	var svgBuilder strings.Builder

	// Start the SVG wrapper. Using a viewBox is crucial for scaling and defining the coordinate system.
	svgBuilder.WriteString(`<svg width="512" height="512" viewBox="0 0 512 512" xmlns="http://www.w3.org/2000/svg">`)
	svgBuilder.WriteString("\n") // Newline for readability

	// Append each layer's content
	for _, layerPath := range layers {
		svgPart, err := g.getRandomSVGPart(layerPath)
		if err != nil {
			// If a layer directory is empty (e.g., no country-specific hair), we skip it.
			if os.IsNotExist(err) {
				continue
			}
			return "", fmt.Errorf("failed to get svg part from layer %s: %w", layerPath, err)
		}
		// Add the SVG content from the layer file
		svgBuilder.WriteString(svgPart)
		svgBuilder.WriteString("\n") // Newline for readability
	}

	svgBuilder.WriteString("</svg>")

	return svgBuilder.String(), nil
}

// collectLayers finds all layer directories based on the specified country and gender.
// It implements a fallback mechanism: country/gender -> common/gender -> country/common -> common/common.
func (g *Generator) collectLayers() ([]string, error) {
	// Use a map to collect unique layer directories, keyed by layer name (e.g., "01_body").
	layerMap := make(map[string]string)

	// Define search paths with decreasing specificity.
	searchPaths := []string{
		filepath.Join(AssetRoot, g.Country, g.Gender),
		filepath.Join(AssetRoot, CommonDir, g.Gender),
		filepath.Join(AssetRoot, g.Country, CommonDir),
		filepath.Join(AssetRoot, CommonDir, CommonDir),
	}

	for _, path := range searchPaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue // Skip if the base directory doesn't exist
		}

		files, err := ioutil.ReadDir(path)
		if err != nil {
			return nil, fmt.Errorf("failed to read directory %s: %w", path, err)
		}

		for _, file := range files {
			if file.IsDir() {
				layerName := file.Name()
				// Only add the layer if it hasn't been found in a more specific path.
				if _, exists := layerMap[layerName]; !exists {
					layerMap[layerName] = filepath.Join(path, layerName)
				}
			}
		}
	}

	// Sort the layers by name to ensure correct drawing order (e.g., 00_background, 01_body, ...).
	var sortedLayers []string
	for _, layerPath := range layerMap {
		sortedLayers = append(sortedLayers, layerPath)
	}
	sort.Strings(sortedLayers)

	return sortedLayers, nil
}

// getRandomSVGPart selects a random SVG from a layer directory, reads it,
// and extracts the content from within the <svg> tags.
func (g *Generator) getRandomSVGPart(layerPath string) (string, error) {
	files, err := ioutil.ReadDir(layerPath)
	if err != nil {
		return "", fmt.Errorf("could not read layer directory %s: %w", layerPath, err)
	}

	var svgs []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".svg") {
			svgs = append(svgs, file.Name())
		}
	}

	if len(svgs) == 0 {
		return "", os.ErrNotExist // Use a specific error to indicate an empty layer.
	}

	randomFile := svgs[rand.Intn(len(svgs))]
	filePath := filepath.Join(layerPath, randomFile)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read svg file %s: %w", filePath, err)
	}

	// Extract the content inside the <svg> tag. This simple string manipulation
	// is sufficient for this project's controlled assets but would be fragile
	// for arbitrary SVG files.
	contentStr := string(content)
	startTagEnd := strings.Index(contentStr, ">")
	endTagStart := strings.LastIndex(contentStr, "</svg>")

	if startTagEnd == -1 || endTagStart == -1 || startTagEnd >= endTagStart {
		return "", fmt.Errorf("could not find valid svg content in %s", filePath)
	}

	// Return the string containing the SVG elements.
	return contentStr[startTagEnd+1 : endTagStart], nil
}