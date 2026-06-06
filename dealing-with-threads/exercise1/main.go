/*
Write a program that accepts a list of text filenames as arguments. 
For each filename, the program should spawn a new goroutine that will output the contents of that file to the console. 
You can use the time.Sleep() function to wait for the child goroutines to complete (until you know how to do this better).
*/

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func getFiles(filenames []string) {
	for _, file := range filenames {

		content, err := os.ReadFile(file)

		if err != nil {
			log.Fatal(err)
		}

		os.Stdout.Write(content)
	}
}

func main() {
	// get time to start
	fmt.Println("Started: ", time.Now().Format("15:04:05"))

	filenames := os.Args[1:]

	if len(filenames) == 0 {
		fmt.Fprintln(os.Stderr, "Please provide at least one filename.")
		os.Exit(1)
	}

	go getFiles(filenames)

	time.Sleep(2 * time.Second)
	fmt.Println("Ended: ", time.Now().Format("15:04:05"))
}
