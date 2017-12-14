package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"bufio"
)

const DateFormat = "2006/01/02 15:04"

func main() {
	tm := time.Now()
	dir, err := os.Getwd()
	if err == nil {
		fmt.Print(dir, "\n")
	}
	fmt.Print("DateTime:", tm.Format(DateFormat), "\n")
	//  Add Git UnStage File
	add, err := exec.Command("git", "add", "--all").CombinedOutput()
	if err != nil {
		log.Print("Error...Add\n")
	}
	//  Commit Git
	scr := bufio.NewScanner(os.Stdin)
	fmt.Print("CommitMessage > ")
	scr.Scan()
	cmt, err := exec.Command("git", "commit", "-m", "Commit :"+tm.Format(DateFormat)+" "+string(scr.Text())).CombinedOutput()
	if err != nil {
		log.Print("Error...Commit\n")
	}
	//  Push Git
	push, err := exec.Command("git", "push").CombinedOutput()
	if err != nil {
		log.Print("Error...Push\n")
	}
	fmt.Print("Git Add", string(add), "\n")
	fmt.Print("Git Commit", string(cmt), "\n")
	fmt.Print("Git Push", string(push), "\n")
}
