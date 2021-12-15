package main

import (
	"fmt"
	"log"

	"nrsDDD/Domain/Models/user"
	userService "nrsDDD/Domain/Services/user"
)

func main() {
	user1, err := user.New("KEY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	err = user1.ChangeName("Kenta")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user1)

	user2, err := user.New("Pori")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user2)

	if user1.Equals(*user2) {
		fmt.Println("Equal")
	} else {
		fmt.Println("Not Equal")
	}

	if userService.Exists(*user2) {
		fmt.Println("Exists")
	} else {
		fmt.Println(("Not Exists"))
	}
}
