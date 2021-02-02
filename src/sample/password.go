package main

import (
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nYour password is %v\n", string(password))
	}
}
