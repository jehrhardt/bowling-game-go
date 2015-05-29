package bowling

type Game struct {
	rolls [21]int
	currentRoll int
}

func (g *Game) Roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll++
}

func (g *Game) Score() int {
	score := 0
	i := 0

	for frame := 0; frame < 10; frame++ {
		score += g.rolls[i]
		score += g.rolls[i + 1]
		i += 2
	}

	return score
}
