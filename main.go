package main

import (
	models "bitva/models"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	ben := models.NewSnake()
	war, err := models.NewWarrior()
	if err != nil {
		logrus.Error("битва не почалась з технічних причин - ", err)
		return
	}
	myCombat := models.NewCombat(*ben, *war)

	fmt.Println(myCombat)
	myCombat.Fight()

}
