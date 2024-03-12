package main

import (
	"fmt"
	"math/rand"
)

func RandGenerator(min, max int) int {
	max++
	return rand.Intn(max-min) + min
}

func HeadsGenerator() int {
	var snakeHeadsGrowVariants []int
	for range 35 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 0)
	}
	for range 35 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 1)
	}
	for range 20 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 2)
	}
	for range 10 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 3)
	}
	return snakeHeadsGrowVariants[RandGenerator(0, 99)]
}

type Warrior struct {
	name string
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

func NewWarrior() *Warrior {
	return &Warrior{
		name: WarriorGetName(),
	}
}

func (w Warrior) WarriorPunch() int {
	return RandGenerator(1, 5)
}

type Combat struct {
	warior Warrior
	snake  Snake
}

func NewCombat(mySnake Snake, myWarrior Warrior) *Combat {
	return &Combat{
		snake:  mySnake,
		warior: myWarrior,
	}
}

func (c Combat) Fight() {
	fmt.Printf("І повстав змій %s із попелу ядерного вогню\n", c.snake.name)
	fmt.Printf("І пішла війна на знищення людства, і йшла десятиліття. На захист людства виступив %s\n", c.warior.name)
	i := 0
	for {
		punch := c.warior.WarriorPunch()
		fmt.Printf("І махнув %s мечем і відлетіло у змія %d голів\n", c.warior.name, punch)
		countNewHeads := c.snake.HeadsGrower(punch)
		fmt.Printf("Змій відрости %d голів\n Всього у змія тепер %d голів\n", countNewHeads, c.snake.headCount)

		if c.snake.headCount == 0 {
			fmt.Println("Змій здох")
			return
		}

		if c.snake.headCount > 200 {
			fmt.Println("Воїн здох")
		}

		i++
		if i == 200 {
			fmt.Println("Всім надоїло і вони пішли спати")
			return
		}
	}

}

func main() {
	ben := NewSnake()
	war := NewWarrior()
	myCombat := NewCombat(*ben, *war)

	println(myCombat.warior.name)
	myCombat.Fight()

}
