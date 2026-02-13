// Package parser for parsing the input file and creating a struct
package parser

import (
	"bufio"
	"os"
	"strings"
)

func Parse(fileName string) (map[string]any, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stru, err := parseForStruct(f)
	if err != nil {
		return nil, err
	}

	return stru, nil
}

func parseForStruct(file *os.File) (map[string]any, error) {
	scanner := bufio.NewScanner(file)
	isStruct := false

	jsonObject := make(map[string]any)

	for scanner.Scan() {
		line := scanner.Text()
		if isStruct && strings.Contains(line, "}") {
			isStruct = false
		}

		if isStruct {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				fieldName := parts[0]
				fieldType := parts[1]
				jsonObject[fieldName] = fieldType
			}
			continue
		}

		if strings.Contains(line, "struct") {
			isStruct = true
			// name := getStructName(line)
			continue
		}
	}

	if err := scanner.Err(); err != nil {
	}

	return jsonObject, nil
}

// func getStructName(line string) string {
// 	name := ""
// 	parts := strings.Fields(line)
// 	if len(parts) > 1 {
// 		name = parts[1]
// 	}
//
// 	return name
// }
