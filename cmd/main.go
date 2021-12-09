package main

import (
	"fmt"
	"log"

	"nrsDDD/user"
)

func main() {
	fullName, err := user.New("Kenta", "Yamaguchi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fullName.FirstName, fullName.LastName)
}
