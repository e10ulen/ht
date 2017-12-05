package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("Test")
	dir, err := os.Getwd()
	if err == nil {
		fmt.Print(dir, "\n")
		//  Add Git UnStage File
		add, err := exec.Command("git", "add", "--all").CombinedOutput()
		if err != nil {
			log.Print("Error...Add")
		}
		//  Commit Git
		cmt, err := exec.Command("git", "commit", "-m", "'update-site'").CombinedOutput()
		if err != nil {
			log.Print("Error...Commit")
		}
		//  Push Git
		push, err := exec.Command("git", "push").CombinedOutput()
		if err != nil {
			log.Print("Error...Push")
		}
		fmt.Print("Git Add", string(add), "\n")
		fmt.Print("Git Commit", string(cmt), "\n")
		fmt.Print("Git Push", string(push), "\n")
	}
}
