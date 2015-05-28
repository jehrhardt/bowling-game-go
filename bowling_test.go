package bowling

import "testing"

func (g *Game) RollMany(runs int, pins int) {
	for i := 0; i < runs; i++ {
		g.Roll(pins)
	}
}

func TestGutterGame(t *testing.T) {
	game := new(Game)
	game.RollMany(20, 0)
	score := game.Score()

	if score != 0 {
		t.Error("Gutter game:", "Expected", 0, "was", score)
	}
}

func TestAllOnes(t *testing.T) {
	game := new(Game)
	game.RollMany(20, 1)
	score := game.Score()

	if score != 20 {
		t.Error("All ones:", "Expected", 20, "was", score)
	}
}
