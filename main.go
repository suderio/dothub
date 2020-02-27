package main

import (
	"flag"
//	"fmt"
	"log"
	"os"
	//	"os/exec"
)

var (
	origin     = flag.String("origin", "master", "Origin (branch, repo or 'local') of the files")
	branch     = flag.String("branch", "$hostname", "Destination branch")
	repository = flag.String("repository", "", "Destination repo")
  file       = flag.String("file", "", "File to add or remove")
)

func main() {
	flag.Parse()
	command := os.Args[1]
	switch command {
  // Initializes the repo (init / clone)
	case "init":
		initialize()
  // Updates (pull / push) the repo
	case "update":
		update()
  // Calls the merge tool
  case "solve":
    solve()
  // Adds a file to the repo
  case "add":
    add(*file)
  // Removes a file from the repo
  case "remove":
    remove(*file)
  // Copy a file to another branch
  case "copy":
    copyToBranch(*file)
  default:
		log.Printf("Command %s unknown\n", command)
	}
}

func initialize() {
	log.Printf("git clone --bare %s\n", *repository)
	log.Printf("git checkout -b %s\n", *branch)
}

func update() {
	log.Printf("git pull")
	log.Printf("git push")
}

func add(filename string) {

}

func remove(filename string) {

}

func solve() {

}

func copyToBranch(filename string) {

}

