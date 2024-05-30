package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("sky.txt") //open the file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var textLines []string
	for scanner.Scan() {
		line := scanner.Text()
		textLines = append(textLines, line)
		//substrings := strings.Split(line, ",")
		//textLines = append(textLines, substrings...)
	}
	fmt.Println("Input text")
	fmt.Println(textLines)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	fmt.Println("Strings that contain the search string:")
	fmt.Println(checkPhrase(textLines, "sky was"))
}

func checkPhrase(textLines []string, phrase string) []string {
	var textLineWhichRepeated []string

	for _, textLine := range textLines {
		if strings.Contains(textLine, phrase) {
			textLineWhichRepeated = append(textLineWhichRepeated, textLine)
		}
	}
	return textLineWhichRepeated
}
