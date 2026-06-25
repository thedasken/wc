package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("veuillez passer le nom du fichier en argument")
		return
	}

	option := "all"
	fileName := ""

	if strings.HasPrefix(os.Args[1], "-") {
		option = os.Args[1]
		if len(os.Args) < 3 {
			fmt.Println("veuillez entrer le nom du fichier à analyser")
			return
		}
		fileName = os.Args[2]
	} else {
		fileName = os.Args[1]
	}

	if !isValidOption(option) {
		fmt.Printf("option %s invalide, options disponibles: -l, -w, -c\n", option)
		return
	}

	lineCount, wordCount, byteCount, err := countFile(fileName)
	if err != nil {
		fmt.Println("erreur:", err)
		return
	}

	switch option {
	case "all":
		fmt.Printf("lines: %d\n", lineCount)
		fmt.Printf("words: %d\n", wordCount)
		fmt.Printf("bytes: %d\n", byteCount)
	case "-l":
		fmt.Printf("lines: %d\n", lineCount)
	case "-w":
		fmt.Printf("words: %d\n", wordCount)
	case "-c":
		fmt.Printf("bytes: %d\n", byteCount)
	}
}

func isValidOption(option string) bool {
	return option == "all" || option == "-l" || option == "-w" || option == "-c"
}

func countFile(fileName string) (int, int, int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, 0, err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return 0, 0, 0, err
	}

	byteCount := info.Size()
	lineCount := 0
	wordCount := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++
		wordCount += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}

	return lineCount, wordCount, byteCount, nil
}
