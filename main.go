package main

import (
	models "bitva/models"
	"fmt"
)

func main() {
	ben := models.NewSnake()
	war := models.NewWarrior()
	myCombat := models.NewCombat(*ben, *war)

	fmt.Println(myCombat)
	myCombat.Fight()

}
