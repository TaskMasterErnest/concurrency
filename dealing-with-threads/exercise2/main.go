/*
Expand the program you wrote in the first exercise so that instead of printing the contents of the text files,
it searches for a string match.
The string to search for is the first argument on the command line.
When you spawn a new goroutine, instead of printing the file’s contents, it should read the file and search for a match.
If the goroutine finds a match,it should output a message saying that the filename contains a match.
*/

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"regexp"
// 	"time"
// )

// func searchFiles(word string, filenames []string) {
// 	for _, file := range filenames {
// 		content, err := os.ReadFile(file)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		re := regexp.MustCompile(word)
// 		if re.Match(content) != false {
// 			fmt.Printf("File: %s, contains a match on the string \"%s\"\n", file, word)
// 		}
// 	}
// }

// func main() {
// 	fmt.Println("Started:", time.Now().Format("15:04:05"))

// 	wordToSearch := os.Args[1]
// 	filenames := os.Args[2:]

// 	if len(wordToSearch) == 0 {
// 		fmt.Fprintln(os.Stderr, "Please enter a word to search.")
// 		os.Exit(1)
// 	}

// 	if len(filenames) == 0 {
// 		fmt.Fprintln(os.Stderr, "Please enter at least one filename.")
// 		os.Exit(1)
// 	}

// 	go searchFiles(wordToSearch, filenames)

// 	time.Sleep(2 * time.Second)
// 	fmt.Println("Ended:", time.Now().Format("15:04:05"))
// }

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"
)

func searchFiles(word string, file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Skipping %s: %v\n", file, err)
		return
	}

	if bytes.Contains(content, []byte(word)) {
		fmt.Printf("File: %s, contains a match on the string \"%s\"\n", file, word)
	}
}

func main() {
	fmt.Println("Started:", time.Now().Format("15:04:05"))

	// check the length of the arguments that will be used
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: go run main.go <word_to_search> <file1> <file2>...")
		os.Exit(1)
	}

	word_to_search := os.Args[1]
	filenames := os.Args[2:]

	for _, file := range filenames {
		go searchFiles(word_to_search, file)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Ended:", time.Now().Format("15:04:05"))
}
