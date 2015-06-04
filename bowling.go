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
	frameIndex := 0

	for frame := 0; frame < 10; frame++ {
		if g.isStrike(frameIndex) {
			score += 10
			score += g.rolls[frameIndex + 1]
			score += g.rolls[frameIndex + 2]
			frameIndex++
		} else if g.isSpare(frameIndex) {
			score += 10
			score += g.rolls[frameIndex + 2]
			frameIndex += 2
		} else {
			score += g.rolls[frameIndex]
			score += g.rolls[frameIndex + 1]
			frameIndex += 2
		}
	}

	return score
}

func (g *Game) isSpare(frameIndex int) bool {
	return g.rolls[frameIndex] + g.rolls[frameIndex + 1] == 10
}

func (g *Game) isStrike(frameIndex int) bool {
	return g.rolls[frameIndex] == 10
}
