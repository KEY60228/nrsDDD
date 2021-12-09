package main

import (
	"fmt"

	"nrsDDD/user"
)

func main() {
	fullName, err := user.New("Kenta", "Yamaguchi")
	if err != nil {
		fmt.Errorf("error!")
	}
	fmt.Println(fullName.FirstName, fullName.LastName)
}
