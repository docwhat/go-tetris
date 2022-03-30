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
