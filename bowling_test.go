package bowling

import "testing"

func TestGutter(t *testing.T) {
	game := new(Game)

	for i := 0; i < 20; i++ {
		game.Roll(0)		
	}

	score := game.Score()

	if score != 0 {
		t.Error("Gutter game:", "Expected", 0, ", was", score)
	}
}
