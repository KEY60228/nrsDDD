package main

import (
	"fmt"
	"log"

	app "nrsDDD/application"
	ur "nrsDDD/infrastructure/pg/user"
	// testur "nrsDDD/infrastructure/testpg/user"
)

func main() {
	userRepository, err := ur.New()
	// userRepository, err := testur.New() // テスト用
	if err != nil {
		log.Fatal(err)
	}

	program, err := app.New(userRepository)
	if err != nil {
		log.Fatal(err)
	}

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

	err = program.DeleteUser(*user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user1")

	err = program.DeleteUser(*user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted user2")
}
