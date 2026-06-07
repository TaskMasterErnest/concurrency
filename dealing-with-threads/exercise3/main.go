/*
Change the program you wrote in the second exercise so that instead of passing a list of text filenames,
you pass a directory path. The program will look inside this directory and list the files.
For each file, you can spawn a goroutine that will search for a string match
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func searchFiles(word, file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	if bytes.Contains(content, []byte(word)) {
		fmt.Printf("File %s, contains a match for the string \"%s\"\n", file, word)
	}
}

func main() {
	fmt.Println("Started:", time.Now().Format("15:04:05"))

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <word_to_search> <path_to_directory>")
		os.Exit(1)
	}

	wordToSearch := os.Args[1]
	dirPath := os.Args[2]

	list, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening named directory.", err)
		os.Exit(1)
	}

	for _, entry := range list {
		// skip if it is a directory
		if entry.IsDir() {
			continue
		}

		// join the directory path with the filename
		fullPath := filepath.Join(dirPath, entry.Name())

		go searchFiles(wordToSearch, fullPath)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Ended:", time.Now().Format("15:04:05"))
}
