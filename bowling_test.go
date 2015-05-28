package bowling

import "testing"

func TestGutterGame(t *testing.T) {
	game := new(Game)

	for i := 0; i < 20; i++ {
		game.Roll(0)		
	}

	score := game.Score()

	if score != 0 {
		t.Error("Gutter game:", "Expected", 0, "was", score)
	}
}

func TestAllOnes(t *testing.T) {
	game := new(Game)

	for i := 0; i < 20; i++ {
		game.Roll(1)
	}

	score := game.Score()

	if score != 20 {
		t.Error("All ones:", "Expected", 20, "was", score)
	}
}
