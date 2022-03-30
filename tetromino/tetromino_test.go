package tetromino_test

import (
	"testing"

	"docwhat.org/go-tetris/tetromino"
)

func Test_IsStringy(t *testing.T) {
	tetrominoes := []tetromino.Tetromino{
		tetromino.I_PIECE, tetromino.J_PIECE, tetromino.L_PIECE, tetromino.O_PIECE, tetromino.S_PIECE, tetromino.T_PIECE, tetromino.Z_PIECE,
	}

	for _, tetromino := range tetrominoes {
		str := tetromino.String()
		if len(str) != 1 {
			t.Errorf("Expected tetromino %d to be a single character, got %q", tetromino, str)
		}
	}
}
