package tetromino

// SRS: https://tetris.fandom.com/wiki/SRS
func (s SRSSprite) Rotate(x, y, clockwise int) int {
	switch len(s.Piece) {
	case 9:
		return s.Rotate9(x, y, clockwise)
	case 12:
		return s.Rotate12(x, y, clockwise)
	case 16:
		return s.Rotate16(x, y, clockwise)
	default:
		panic("unreachable")
	}
}

func (s SRSSprite) Rotate9(x, y, clockwise int) int {
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
		return 8 - y*3 - x
	case 3:
		// 2 5 8
		// 1 4 7
		// 0 3 6
		return 2 - y + x*3
	default:
		panic("invalid clockwise")
	}
}

func (s SRSSprite) Rotate12(x, y, clockwise int) int {
	// Never rotates.
	// 0 1 2 3
	// 4 5 6 7
	// 8 9 A B
	return y*4 + x
}

func (s SRSSprite) Rotate16(x, y, clockwise int) int {
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

type SRSSprite struct {
	Width  int
	Height int
	Piece  []bool
}

func NewSRSSprite(mask string) SRSSprite {
	size := len(mask)
	sprite := SRSSprite{}
	switch size {
	case 9:
		sprite.Width = 3
		sprite.Height = 3
	case 12:
		sprite.Width = 4
		sprite.Height = 3
	case 16:
		sprite.Width = 4
		sprite.Height = 4
	default:
		panic("invalid mask size")
	}

	sprite.Piece = make([]bool, size)
	for i := 0; i < size; i++ {
		sprite.Piece[i] = mask[i] != '.'
	}

	return sprite
}

var (
	SRS_I_PIECE = NewSRSSprite("....IIII........")
	SRS_J_PEICE = NewSRSSprite("J..JJJ...")
	SRS_L_PEICE = NewSRSSprite("..LLLL...")
	SRS_O_PIECE = NewSRSSprite(".OO..OO.....")
	SRS_S_PEICE = NewSRSSprite(".SSSS....")
	SRS_T_PEICE = NewSRSSprite(".T.TTT...")
	SRS_Z_PEICE = NewSRSSprite("ZZ..ZZ...")
)
