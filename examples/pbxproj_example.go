// Example program demonstrating pbxproj comment preservation
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"howett.net/plist"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: pbxproj_example <pbxproj-file>")
		fmt.Println("Example: pbxproj_example project.pbxproj")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Read the pbxproj file
	fmt.Printf("Reading %s...\n", inputFile)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Method 1: Standard plist (loses comments)
	fmt.Println("\n=== Method 1: Standard plist (loses comments) ===")
	var standardResult map[string]interface{}
	_, err = plist.Unmarshal(data, &standardResult)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		os.Exit(1)
	}
	standardOutput, _ := plist.Marshal(standardResult, plist.OpenStepFormat)
	standardComments := countComments(string(standardOutput))
	fmt.Printf("Comments in standard output: %d\n", standardComments)

	// Method 2: Comment-preserving parser
	fmt.Println("\n=== Method 2: Comment-preserving parser ===")
	dict, err := plist.ParsePbxProj(string(data))
	if err != nil {
		fmt.Printf("Error parsing with comment preservation: %v\n", err)
		os.Exit(1)
	}
	preservedOutput := plist.GeneratePbxProj(dict)
	preservedComments := countComments(preservedOutput)
	fmt.Printf("Comments in preserved output: %d\n", preservedComments)

	// Show statistics
	originalComments := countComments(string(data))
	fmt.Printf("\n=== Statistics ===\n")
	fmt.Printf("Original file comments: %d\n", originalComments)
	fmt.Printf("Standard method comments: %d (lost %d)\n", standardComments, originalComments-standardComments)
	fmt.Printf("Preserved method comments: %d (lost %d)\n", preservedComments, originalComments-preservedComments)

	// Write both outputs for comparison
	outputDir := filepath.Dir(inputFile)
	baseName := filepath.Base(inputFile)

	standardFile := filepath.Join(outputDir, baseName+".standard")
	preservedFile := filepath.Join(outputDir, baseName+".preserved")

	os.WriteFile(standardFile, standardOutput, 0644)
	os.WriteFile(preservedFile, []byte(preservedOutput), 0644)

	fmt.Printf("\nOutput files written:\n")
	fmt.Printf("  Standard (no comments): %s\n", standardFile)
	fmt.Printf("  Preserved (with comments): %s\n", preservedFile)
	fmt.Printf("\nYou can compare them with: diff %s %s\n", standardFile, preservedFile)
}
