package tetromino_test

import (
	"testing"

	"docwhat.org/go-tetris/tetromino"
)

func Test_I_PeiceRotates180(t *testing.T) {
	piece := tetromino.SRS_I_PIECE
	clockwise := uint(2)
	expected := uint(0)
	actual := piece.Rotate(0, 0, clockwise)
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
