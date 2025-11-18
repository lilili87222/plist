// Example: Modify pbxproj while preserving comments
package main

import (
	"fmt"
	"os"

	"github.com/lilili87222/plist"
)

func main() {
	// Sample pbxproj content with comments
	pbxprojContent := `// !$*UTF8*$!
{
	archiveVersion = 1;
	classes = {
	};
	objectVersion = 56;
	objects = {
		ABC123 /* MyApp */ = {
			isa = PBXProject;
			buildConfigurationList = DEF456 /* Build configuration list for PBXProject "MyApp" */;
			targets = (
				TARGET1 /* MyApp */,
				TARGET2 /* MyAppTests */,
			);
		};
		DEF456 /* Build configuration list for PBXProject "MyApp" */ = {
			isa = XCConfigurationList;
			buildConfigurations = (
				CONFIG1 /* Debug */,
				CONFIG2 /* Release */,
			);
		};
	};
	rootObject = ABC123 /* MyApp */;
}
`

	fmt.Println("=== Original Content ===")
	fmt.Println(pbxprojContent)

	// Parse with comment preservation
	dict, err := plist.ParsePbxProj(pbxprojContent)
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n=== Parsed Structure ===\n")
	fmt.Printf("Root keys: %v\n", dict.Keys)
	fmt.Printf("Number of root keys: %d\n", len(dict.Keys))

	// Demonstrate accessing nested values
	for i, key := range dict.Keys {
		value := dict.Values[i]
		comment := dict.Comments[key]

		fmt.Printf("\nKey: %s", key)
		if comment != "" {
			fmt.Printf(" /* %s */", comment)
		}
		fmt.Println()

		// Show value type
		switch v := value.Value.(type) {
		case *plist.PbxProjDict:
			fmt.Printf("  Type: Dictionary with %d keys\n", len(v.Keys))
			if key == "objects" {
				// Show objects
				for j, objKey := range v.Keys {
					objComment := v.Comments[objKey]
					fmt.Printf("    - %s", objKey)
					if objComment != "" {
						fmt.Printf(" /* %s */", objComment)
					}
					fmt.Println()

					// Show nested dictionary info
					if objDict, ok := v.Values[j].Value.(*plist.PbxProjDict); ok {
						for k, nestedKey := range objDict.Keys {
							if nestedKey == "isa" {
								if isaValue, ok := objDict.Values[k].Value.(string); ok {
									fmt.Printf("      isa: %s\n", isaValue)
								}
							}
						}
					}
				}
			}
		case *plist.PbxProjArray:
			fmt.Printf("  Type: Array with %d elements\n", len(v.Values))
		case string:
			fmt.Printf("  Type: String, Value: %s\n", v)
		}
	}

	// Example: Add a new object while preserving comments
	fmt.Println("\n=== Modifying Content ===")

	// Find the objects dictionary
	for i, key := range dict.Keys {
		if key == "objects" {
			if objectsDict, ok := dict.Values[i].Value.(*plist.PbxProjDict); ok {
				// Add a new object
				newKey := "NEWOBJ123"
				newValue := &plist.PbxProjValue{
					Value: &plist.PbxProjDict{
						Keys: []string{"isa", "name"},
						Values: []*plist.PbxProjValue{
							{Value: "PBXFileReference", Comment: ""},
							{Value: "NewFile.swift", Comment: ""},
						},
						Comments: make(map[string]string),
					},
					Comment: "",
				}

				objectsDict.Keys = append(objectsDict.Keys, newKey)
				objectsDict.Values = append(objectsDict.Values, newValue)
				objectsDict.Comments[newKey] = "New file added programmatically"

				fmt.Printf("Added new object: %s /* %s */\n", newKey, objectsDict.Comments[newKey])
			}
		}
	}

	// Generate output with all comments preserved
	output := plist.GeneratePbxProj(dict)

	fmt.Println("\n=== Modified Content (Comments Preserved) ===")
	fmt.Println(output)

	// Count comments
	originalComments := countComments(pbxprojContent)
	outputComments := countComments(output)

	fmt.Printf("\n=== Comment Statistics ===\n")
	fmt.Printf("Original comments: %d\n", originalComments)
	fmt.Printf("Output comments: %d (including 1 new)\n", outputComments)
	fmt.Printf("Comments preserved: %d out of %d\n", outputComments-1, originalComments)
}

func countComments(s string) int {
	count := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '/' && s[i+1] == '*' {
			count++
		}
	}
	return count
}
