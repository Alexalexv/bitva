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
		newHeads += HeadsGenerator()
	}
	return newHeads

}
