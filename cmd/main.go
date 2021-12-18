package main

import (
	"fmt"
	"log"

	app "nrsDDD/application"
)

func main() {
	program := app.New() // 仮
	user1, err := program.CreateUser("Kenta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	user2, err := program.CreateUser("Pori")
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
