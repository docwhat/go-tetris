package tgm3

import (
	"fmt"
	"testing"
)

func Test_NewRandomizer(t *testing.T) {
	expected := 7 * 5
	randomizer := NewRandomizer()
	if len(randomizer.Tetrominos) != expected {
		t.Errorf("Expected %d pieces, got %d", expected, len(randomizer.Tetrominos))
	}
}

func Test_FirstPiece(t *testing.T) {
	max := 100
	expected := HistoryStack{I_PIECE, J_PIECE, L_PIECE, T_PIECE}
	for i := 0; i < max; i++ {
		t.Run(fmt.Sprintf("check %d of %d", i, max), func(t *testing.T) {
			randomizer := NewRandomizer()
			piece := randomizer.Next()
			if !expected.Contains(piece) {
				t.Errorf("Expected first piece to be one of %v, got %v", expected, piece)
			}
		})
	}
}
