package main

import (
	"fmt"

	"github.com/iiqbalt/my-first-go-app/models"
	// "my-first-gomodels"
)

func main() {
	fmt.Println("my first go app")

	// name := models.GetCity("sidoarjo")
	name := models.Address{
		Name:    "fatikj",
		City:    "sidoarjo",
		Country: "",
	}

	fmt.Println(name.SayAddress())

	fmt.Println(models.GetName(&name))
}
