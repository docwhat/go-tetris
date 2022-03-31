package tetromino

// SRS: https://tetris.fandom.com/wiki/SRS
type SRS3x3Sprite struct {
	Piece [9]bool
}

func (s SRS3x3Sprite) Rotate(x, y, clockwise uint) uint {
	switch clockwise % 4 {
	case 0:
		// 0 1 2
		// 3 4 5
		// 6 7 8
		return y*3 + x
	case 1:
		// 6 3 0
		// 7 4 1
		// 8 5 2
		return 6 + y - x*3
	case 2:
		// 8 7 6
		// 5 4 3
		// 2 1 0
		return 8 - y - x*3
	case 3:
		// 2 5 8
		// 1 4 7
		// 0 3 6
		return 2 - y + x*3
	default:
		panic("invalid clockwise")
	}
}

func (s SRS4x3Sprite) Rotate(x, y, clockwise uint) uint {
	// Never rotates.
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 A B
	return y*4 + x
}

func (s SRS4x4Sprite) Rotate(x, y, clockwise uint) uint {
	switch clockwise % 4 {
	case 0:
		// 0 1 2 3
		// 4 5 6 7
		// 8 9 A B
		// C D E F
		return y*4 + x
	case 1:
		// C 8 4 0
		// D 9 5 1
		// E A 6 2
		// F B 7 3
		return 12 + y - x*4
	case 2:
		// F E D C
		// B A 9 8
		// 7 6 5 4
		// 3 2 1 0
		return 15 - y - x*4
	case 3:
		// 3 7 B F
		// 2 6 A E
		// 1 5 9 D
		// 0 4 8 C
		return 3 - y + x*4
	default:
		panic("invalid clockwise")
	}
}

type SRS4x4Sprite struct {
	Piece [16]bool
}

type SRS4x3Sprite struct {
	Piece [12]bool
}

func NewSRS3x3Sprite(mask string) SRS3x3Sprite {
	var s [9]bool
	for i := 0; i < len(s); i++ {
		s[i] = mask[i] != '.'
	}
	return SRS3x3Sprite{s}
}

func NewSRS4x3Sprite(mask string) SRS4x3Sprite {
	var s [12]bool
	for i := 0; i < len(s); i++ {
		s[i] = mask[i] != '.'
	}
	return SRS4x3Sprite{s}
}

func NewSRS4x4Sprite(mask string) SRS4x4Sprite {
	var s [16]bool
	for i := 0; i < len(s); i++ {
		s[i] = mask[i] != '.'
	}
	return SRS4x4Sprite{s}
}

var (
	SRS_I_PIECE = NewSRS4x4Sprite("....IIII........")
	SRS_J_PEICE = NewSRS3x3Sprite("J..JJJ...")
	SRS_L_PEICE = NewSRS3x3Sprite("..LLLL...")
	SRS_O_PIECE = NewSRS4x3Sprite(".OO..OO.....")
	SRS_S_PEICE = NewSRS3x3Sprite(".SSSS....")
	SRS_T_PEICE = NewSRS3x3Sprite(".T.TTT...")
	SRS_Z_PEICE = NewSRS3x3Sprite("ZZ..ZZ...")
)
