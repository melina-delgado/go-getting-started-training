package main

import (
	"fmt"

	"github.com/melina-delgado/go-getting-started-training/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan",
	}
	fmt.Println(u)
}
