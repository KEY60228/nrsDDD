package main

import (
	"fmt"
	"log"

	uas "nrsDDD/application/user"
	us "nrsDDD/domain/services/user"
	ur "nrsDDD/infrastructure/pg/user"
	// testur "nrsDDD/infrastructure/testpg/user" // テスト用
)

func main() {
	userRepository, err := ur.New()
	// userRepository, err := testur.New() // テスト用
	if err != nil {
		log.Fatal(err)
	}

	userService, err := us.New(userRepository)
	if err != nil {
		log.Fatal(err)
	}

	userApplicationService, err := uas.New(userRepository, *userService)
	if err != nil {
		log.Fatal(err)
	}

	// err = userApplicationService.Register("Kenta")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = userApplicationService.Register("Pori")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	user, err := userApplicationService.Get("3df7db6d-bc20-4591-839a-b6c2581011b8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user)
}
