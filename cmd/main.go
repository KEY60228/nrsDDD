package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

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

	userRegisterService, err := uas.NewUserRegisterService(userRepository, *userService)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Input user name")
		fmt.Print("> ")

		sc := bufio.NewScanner(os.Stdin)

		var name string
		if sc.Scan() {
			name = sc.Text()
		}

		userRegisterCommand, err := uas.NewUserRegisterCommand(name)
		if err != nil {
			log.Fatal(err)
		}

		err = userRegisterService.Handle(*userRegisterCommand)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("--------")
		fmt.Printf("user %v created\n", name)
		fmt.Println("--------")

		var flg bool
		for {
			fmt.Println("continue? (y/N)")
			fmt.Print("> ")

			var yesOrNo string
			if sc.Scan() {
				yesOrNo = sc.Text()
			}

			if yesOrNo == "y" {
				break
			} else if yesOrNo == "N" {
				flg = true
				break
			}

			fmt.Println("invalid command")
		}
		if flg {
			break
		}
	}
}
