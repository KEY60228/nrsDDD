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

	user, err := userApplicationService.Get("823fc68e-36cd-4416-be1b-c96c53983277")
	if err != nil {
		log.Fatal(err)
	}
	if user == nil {
		log.Fatal("ユーザーが存在しません")
	}
	fmt.Println(*user)

	updateCommand, err := uas.NewUserUpdateCommand(user.Id)
	if err != nil {
		log.Fatal(err)
	}

	err = updateCommand.SetName("Ken")
	if err != nil {
		log.Fatal(err)
	}

	user, err = userApplicationService.Update(*updateCommand)
	if err != nil {
		log.Fatal(err)
	}

	deleteCommand, err := uas.NewUserDeleteCommand(user.Id)
	err = userApplicationService.Delete(*deleteCommand)

	fmt.Println(*user)
}
