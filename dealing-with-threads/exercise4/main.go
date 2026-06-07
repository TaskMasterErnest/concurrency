/*
Adapt the program in the third exercise to continue searching recursively in any subdirectories.
If you give your search goroutine a file, it should search for a string match in that file,
just like in the previous exercises.
Otherwise, if you give it a directory, it should recursively spawn a new goroutine for each file or directory found inside.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
)

func searchFiles(word, path string) {

	// checking if it is a file
	info, err := os.Stat(path)
	if err != nil {
		log.Printf("Skipping %s: %v", path, err)
		return
	}

	// checking if it is a directory
	if info.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("Error reading current directory %s: %v", path, err)
			return
		}

		for _, entry := range entries {
			fullPath := filepath.Join(path, entry.Name())

			// spawn goroutines to search these files
			go searchFiles(word, fullPath)
		}

		return
	}

	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Skipping %s: %v", path, err)
		return
	}

	if bytes.Contains(content, []byte(word)) {
		// adding color
		pathColour := color.New(color.FgYellow, color.Italic, color.Bold)
		wordColour := color.New(color.FgGreen, color.Italic, color.Bold)

		fmt.Printf("File %s, contains a match for string \"%s\"\n", pathColour.Sprint(path), wordColour.Sprint(word))
	}
}

func main() {
	fmt.Println("Started:", time.Now().Format("15:04:05"))

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: go run main.go <word_to_search> <dir> <file>...")
		os.Exit(1)
	}

	wordToSearch := os.Args[1]
	filePath := os.Args[2]

	go searchFiles(wordToSearch, filePath)

	time.Sleep(3 * time.Second)
	fmt.Println("Ended:", time.Now().Format("15:04:05"))
}
