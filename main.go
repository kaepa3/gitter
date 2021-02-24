package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
)

// Basic example of how to clone a repository using clone options.
func main() {
	directory := os.Args[1]

	// Clone the given repository to the given directory
	r, err := git.PlainOpen(directory)

	CheckIfError(err)

	ref, err := r.Head()
	// ... retrieving the branch being pointed by HEAD
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	var messages []string
	for {
		cmt, err := cIter.Next()
		if err != nil {
			break
		}
		messages = append(messages, cmt.Message)
	}
	fmt.Println(messages)

}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
