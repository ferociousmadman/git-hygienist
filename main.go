package main

import (
	"fmt"
	"strings"
)

type Commit struct {
	Hash    string
	Author  string
	Message string
}

var (
	CommitOne = Commit{Hash: "a1b2c3d", Author: "ferociousmadman", Message: "feat: initial commit"}
	CommitTwo = Commit{Hash: "e5f6g7h", Author: "ferociousmadman", Message: "fix: bug in main.go"}
)

func main() {
	fmt.Printf("Git-Hygienist: Checking your commit health...\n")
	commitObjects := []Commit{CommitOne, CommitTwo}
	// fmt.Printf("%v\n", commitObjects)

	for i := range len(commitObjects) {
		fmt.Printf("%v\n", commitObjects[i])
		if strings.HasPrefix(commitObjects[i].Message, "feat:") {
			println("commit type feat")
		} else if strings.HasPrefix(commitObjects[i].Message, "fix:") {
			println("commit type fix")
		} else if strings.HasPrefix(commitObjects[i].Message, "docs:") {
			println("commit type docs")
		} else {
			println("fail")
		}
		// println("i has printed this many times %v", i)
	}

	// return
}
