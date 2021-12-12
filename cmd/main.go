package main

import (
	"fmt"
	"log"

	"nrsDDD/money"
	"nrsDDD/name"
	"nrsDDD/user"
)

func main() {
	fullName, err := name.New("Kenta", "Yamaguchi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullName.FirstName, fullName.LastName)

	mon, err := money.New(200, "yen")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*mon)
	mon2, err := money.New(300, "yen")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*mon2)
	mon3, err := mon.Add(*mon2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*mon3)

	user, err := user.New("KEY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*user)
}
