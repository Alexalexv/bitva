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
	percentage := RandGenerator(1, 100)
	if percentage > 35 && percentage <= 70 {
		return 1
	}
	if percentage > 70 && percentage <= 90 {
		return 2
	}
	if percentage > 90 {
		return 3
	}
	return 0
}
