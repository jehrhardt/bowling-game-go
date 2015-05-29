package bowling

type Game struct {
	rolls [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.currentRoll++
	g.rolls[g.currentRoll] = pins
}

func (g *Game) Score() int {
	score := 0

	for _, roll := range g.rolls {
		score += roll
	}

	return score
}
