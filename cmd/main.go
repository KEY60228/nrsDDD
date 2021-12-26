package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	uas "nrsDDD/application/user"
	us "nrsDDD/domain/services/user"
	ur "nrsDDD/infrastructure/pg/user"
	// testur "nrsDDD/infrastructure/testpg/user" // テスト用
)

func main() {
	dsn := fmt.Sprintf("host=pgsql dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))
	// dsn := fmt.Sprintf("host=testpg dbname=nrsDDD user=%s password=%s sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD")) // テスト用
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	userRepository, err := ur.New(db)
	// userRepository, err := testur.New(db) // テスト用
	if err != nil {
		log.Fatal(err)
	}

	userService, err := us.New(userRepository)
	if err != nil {
		log.Fatal(err)
	}

	userRegisterService, err := uas.NewUserRegisterService(userRepository, *userService, &ur.UserFactory{})
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
