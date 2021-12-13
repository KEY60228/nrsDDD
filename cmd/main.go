package main

import (
	"fmt"
	"log"

	"nrsDDD/user"
)

func main() {
	user1, err := user.New("1", "KEY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	err = user1.ChangeName("Kenta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	user2, err := user.New("2", "Pori")
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
