package models

type Snake struct {
	name      string
	headCount int
}

func NewSnake() *Snake {
	return &Snake{
		name:      "Горинич",
		headCount: RandGenerator(50, 150),
	}
}

func (s *Snake) HeadsGrower(fallenHeadsCount int) int {
	var newHeads int
	s.headCount -= fallenHeadsCount
	if fallenHeadsCount > s.headCount {
		return 0
	}
	for range fallenHeadsCount {
		newHeads += s.HeadsGenerator()
	}
	return newHeads

}

func (s Snake) HeadsGenerator() int {
	var snakeHeadsGrowVariants []int
	for range 10 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 3)
	}
	for range 20 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 2)
	}
	for range 35 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 1)
	}
	for range 35 {
		snakeHeadsGrowVariants = append(snakeHeadsGrowVariants, 0)
	}

	return snakeHeadsGrowVariants[RandGenerator(0, 99)]
}
