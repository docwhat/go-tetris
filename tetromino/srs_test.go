package tetromino_test

import (
	"fmt"
	"testing"

	"docwhat.org/go-tetris/tetromino"
)

// func Test_I_PeiceRotates180(t *testing.T) {
// 	piece := tetromino.SRS_I_PIECE
// 	clockwise := uint(2)
// 	expected := uint(0)
// 	actual := piece.Rotate(0, 0, clockwise)
// 	if expected != actual {
// 		t.Errorf("Expected %v, got %v", expected, actual)
// 	}
// }

type FallingPiece struct {
	x      int
	y      int
	sprite tetromino.SRSSprite
}

func Test_Silly(t *testing.T) {
	screenWidth := 80
	screenHeight := 30
	fieldWidth := 12
	fieldHeight := 18
	screen := make([]string, screenWidth*screenHeight)

	currentPiece := FallingPiece{
		x:      0,
		y:      0,
		sprite: tetromino.SRS_S_PEICE,
	}

	for x := 0; x < fieldWidth; x++ {
		for y := 0; y < fieldHeight; y++ {
			screen[(x+2)+(y+2)*screenWidth] = "."
		}
	}

	for px := 0; px < currentPiece.sprite.Width; px++ {
		for py := 0; py < currentPiece.sprite.Height; py++ {
			if currentPiece.sprite.Piece[currentPiece.sprite.Rotate(px, py, 3)] {
				screen[(currentPiece.y+py+2)*screenWidth+(currentPiece.x+px+2)] = "X"
			}
		}
	}

	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			fmt.Print(screen[y*screenWidth+x])
		}
		fmt.Print("\n")
	}

	t.Error("boom")
}
