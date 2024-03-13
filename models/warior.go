package models

import (
	"errors"
	"fmt"
	"strings"
)

type Warrior struct {
	name string
}

func NewWarrior() (*Warrior, error) {
	name, err := WarriorGetName()
	if err != nil {
		return nil, err
	}
	return &Warrior{
		name: name,
	}, nil
}

func WarriorGetName() (string, error) {
	var (
		WariorNameVariants    = [3]string{"Олег", "Чорний", "Віктор"}
		WariorSurnameVariants = [3]string{"Винник", "Лицар", "Федорович"}
		WariorDignityVariants = [3]string{"Гучний", "Оболонський", "Слабкий"}
	)
	name := WariorNameVariants[RandGenerator(0, 2)]
	surname := WariorSurnameVariants[RandGenerator(0, 2)]
	dignity := WariorDignityVariants[RandGenerator(0, 2)]
	fullName := fmt.Sprintf("%s %s %s", name, surname, dignity)
	if strings.Contains(fullName, "Слабкий") {
		return fullName, errors.New("воїн не прийшов")
	}
	return fullName, nil

}

func (w Warrior) WarriorPunch() int {
	return RandGenerator(1, 5)
}
