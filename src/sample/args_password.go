package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	flag.Parse()
	//args := flag.Args()
	fmt.Println(flag.Arg(0))

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nYour password is %v\n", string(password))
	}
}
