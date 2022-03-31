package tetromino

type Tetromino uint

const (
	I_PIECE Tetromino = iota
	J_PIECE
	L_PIECE
	O_PIECE
	S_PIECE
	T_PIECE
	Z_PIECE
)

func (t Tetromino) String() string {
	switch t {
	case I_PIECE:
		return "I"
	case J_PIECE:
		return "J"
	case L_PIECE:
		return "L"
	case O_PIECE:
		return "O"
	case S_PIECE:
		return "S"
	case T_PIECE:
		return "T"
	case Z_PIECE:
		return "Z"
	default:
		return "?"
	}
}

func (t Tetromino) GameBoyBlocks(clockwise uint) string {
	switch t {
	case I_PIECE:
		switch clockwise % 2 {
		case 0:
			return "...." +
				"...." +
				"xxxx" +
				"...."
		case 1:
			return ".x.." +
				".x.." +
				".x.." +
				".x.."
		}
	case O_PIECE:
		return "...." +
			".xx." +
			".xx."
		"...."
	case J_PIECE:
		switch clockwise % 4 {
		case 0:
			return "..." +
				"xxx" +
				"..x"
		case 1:
			return ".x." +
				".x." +
				"xx."
		case 2:
			return "x.." +
				"xxx" +
				"..."
		case 3:
			return ".xx" +
				".x." +
				".x."
		}
	case L_PIECE:
		switch clockwise % 4 {
		case 0:
			return "..." +
				"xxx" +
				"x.."
		case 1:
			return "xx." +
				".x." +
				".x."
		case 2:
			return "..x" +
				"xxx" +
				"..."
		case 3:
			return ".x." +
				".x." +
				".xx"
		}
	case S_PIECE:
		switch clockwise % 2 {
		case 0:
			return "..." +
				".xx" +
				"xx."
		case 1:
			return "x.." +
				"xx." +
				".x."
		}
	case T_PIECE:
		switch clockwise % 4 {
		case 0:
			return "..." +
				"xxx" +
				".x."
		case 1:
			return ".x." +
				"xx." +
				".x."
		case 2:
			return ".x." +
				"xxx" +
				"..."
		case 3:
			return ".x." +
				".xx" +
				".x."
		}
	case Z_PIECE:
		switch clockwise % 2 {
		case 0:
			return "..." +
				"xx." +
				".xx"
		case 1:
			return ".x." +
				"xx." +
				"x.."
		}

	}

}
