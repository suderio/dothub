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
		log.Printf("Command %s unknown\n", command)
	}
}

func initialize() {
	fmt.Printf("git clone --bare %s\n", *repository)
	fmt.Printf("git checkout -b %s\n", *branch)
}

func update() {
	fmt.Printf("git pull")
	fmt.Printf("git push")
}

func add(filename string) {

}

func remove(filename string) {

}
