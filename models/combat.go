package models

import (
	"fmt"
)

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
	for i := 0; ; i++ {
		punch := c.warior.WarriorPunch()

		if punch < c.snake.headCount {
			fmt.Printf("І махнув %s мечем і відлетіло у змія %d голів\n", c.warior.name, punch)
			countNewHeads := c.snake.HeadsGrower(punch)
			fmt.Printf("Змій відростив %d голів\nВсього у змія тепер %d голів\n", countNewHeads, c.snake.headCount)
		}

		if punch >= c.snake.headCount {
			fmt.Printf("Точним ударом %s відтяв все що залишилось, змій повністю втратив мотивацію до битви\n", c.warior.name)
			return
		}

		if c.snake.headCount > 200 {
			panic("витязь запанікував, кинув меча та побіг кудись")
		}
	}

}
