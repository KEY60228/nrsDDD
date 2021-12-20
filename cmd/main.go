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

	userRegisterService, err := uas.NewUserRegisterService(userRepository, *userService)
	if err != nil {
		log.Fatal(err)
	}

	userRegisterCommand, err := uas.NewUserRegisterCommand("Kenta")
	if err != nil {
		log.Fatal(err)
	}

	err = userRegisterService.Handle(*userRegisterCommand)
	if err != nil {
		log.Fatal(err)
	}

	user, err := userApplicationService.Get("8752be00-0611-4846-b02a-87dc5b330d98")
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

	userDeleteService, err := uas.NewUserDeleteService(userRepository)
	if err != nil {
		log.Fatal(err)
	}

	deleteCommand, err := uas.NewUserDeleteCommand(user.Id)
	if err != nil {
		log.Fatal(err)
	}

	err = userDeleteService.Handle(*deleteCommand)

	fmt.Println(*user)
}
