package models

import "fmt"

type Warrior struct {
	name string
}

func NewWarrior() *Warrior {
	return &Warrior{
		name: WarriorGetName(),
	}
}

func WarriorGetName() string {
	var (
		WariorNameVariants    = [3]string{"Name_1", "Name_2", "Name_3"}
		WariorSurnameVariants = [3]string{"Surname_1", "Surname_2", "Surname_3"}
		WariorDignityVariants = [3]string{"Dignity_1", "Dignity_2", "Dignity_3"}
	)
	name := WariorNameVariants[RandGenerator(0, 2)]
	surname := WariorSurnameVariants[RandGenerator(0, 2)]
	dignity := WariorDignityVariants[RandGenerator(0, 2)]
	return fmt.Sprintf("%s %s %s", name, surname, dignity)

}

func (w Warrior) WarriorPunch() int {
	return RandGenerator(1, 5)
}
