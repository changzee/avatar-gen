# AGENTS.md - SVG Avatar Generator

Welcome, agent! This document provides instructions for working on the Go-based **SVG** avatar generator project.

## Project Overview

The goal of this project is to generate unique, randomized cartoon avatars in **SVG format** based on user-provided inputs for `country` and `gender`. The generation process works by combining the content of multiple SVG layer files into a single, final SVG file.

## Codebase Structure

-   `cmd/avatar-generator/main.go`: This is the main entry point of the application. It handles command-line flag parsing (`--country`, `--gender`, `--output`), calls the avatar generator, and saves the resulting **SVG string** to a `.svg` file.
-   `pkg/avatar/generator.go`: This is the core logic of the project. The `Generator` struct is responsible for:
    1.  Discovering the available SVG layer directories based on a fallback system (`country/gender` -> `common/gender` -> `country/common` -> `common/common`).
    2.  Sorting the layers numerically to ensure correct stacking (e.g., background first, then body, then clothes).
    3.  Randomly selecting one SVG file from each layer directory.
    4.  Reading the content of each selected SVG file, extracting the elements, and concatenating them into a single SVG string.
-   `assets/`: This directory contains all the **SVG** resources. It is crucial to maintain the directory structure. All assets should be valid SVG files.
-   `output/`: This is the default directory where generated avatars are saved.

## How to Work with the Code

### 1. Running the Application
To test the full functionality, run the `main` package with the required flags.

```bash
# Run the generator for a Chinese male SVG avatar
go run ./cmd/avatar-generator --country=china --gender=male

# Run the generator for a USA female SVG avatar
go run ./cmd/avatar-generator --country=usa --gender=female
```
Verify that a new **`.svg`** file is created in the `output/` directory. You can open this file in a web browser or vector graphics editor to inspect it.

### 2. Modifying the Generator Logic
If you need to change how avatars are generated, `pkg/avatar/generator.go` is the file to edit. The current implementation reads SVG files as plain text and extracts the content within the `<svg>` tags. This is a simple but effective method for controlled assets. For more complex SVG manipulation, you might consider using an XML parsing library.

### 3. Adding or Changing Assets
-   **Asset format must be SVG.**
-   **Use a consistent `viewBox`**. To ensure layers align correctly, all source SVG files should share the same `viewBox` (e.g., `viewBox="0 0 512 512"`). This creates a common coordinate system.
-   **Asset paths are important.** Follow the `assets/{country}/{gender}/{layer_name}` structure. Layer folders must be prefixed with a two-digit number to control the drawing order (e.g., `00_background`, `01_body`).

### 4. Verification
Before submitting any changes, ensure the program compiles and runs without errors.
-   Run `go build ./...` to check for compilation errors.
-   Run the generator with a few different flag combinations to ensure it behaves as expected.
-   Confirm that the output `.svg` file is created, is not empty, and renders correctly in a viewer.

This project's logic is now based on string and file manipulation. Adhering to the established asset structure is key.