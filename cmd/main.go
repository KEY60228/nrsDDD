package main

import (
	"fmt"
	"log"

	application "nrsDDD/Application"
)

func main() {
	p := &application.Program{}
	user1, err := p.CreateUser("Kenta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	user2, err := p.CreateUser("Kenta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user2)

	if user1.Equals(*user2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}
}
