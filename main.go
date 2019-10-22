package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	//	"os/exec"
)

var (
	origin     = flag.String("origin", "master", "Origin (branch, repo or 'local') of the files")
	branch     = flag.String("branch", "$hostname", "Destination branch")
	repository = flag.String("repository", "", "Destination repo")
)

func main() {
	flag.Parse()
	command := os.Args[1]
	switch command {
	case "init":
		initialize()
	case "update":
		update()
	default:
		log.Printf("Command %q unknown", command)
	}
}

func initialize() {
	fmt.Printf("git clone --bare %q", repository)
	fmt.Printf("git checkout -b %q", branch)
}

func update() {
	fmt.Printf("git pull")
	fmt.Printf("git push")
}

func add(filename string) {

}

func remove(filename string) {

}
