package tgm3

import "math/rand"

// https://simon.lc/the-history-of-tetris-randomizers

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

// A HistoryStack is used to ensure that rolled pieces are not repeated (flooded) too often.
type HistoryStack []Tetromino

func (stack *HistoryStack) Push(piece Tetromino) {
	if len(*stack) >= 4 {
		*stack = (*stack)[1:]
	}
	*stack = append((*stack), piece)
}

func (stack *HistoryStack) Contains(piece Tetromino) bool {
	for _, val := range *stack {
		if val == piece {
			return true
		}
	}
	return false
}

// An OrderStack is used to prevent a lack of a particular piece (drought).
type OrderStack []Tetromino

func (stack *OrderStack) Peek() Tetromino {
	return (*stack)[0]
}

func (stack *OrderStack) Push(piece Tetromino) {
	for idx, val := range *stack {
		if val == piece {
			*stack = append((*stack)[:idx], (*stack)[idx+1:]...)
		}
	}
	*stack = append(*stack, piece)
}

type Randomizer struct {
	Tetrominos []Tetromino
	history    HistoryStack
	order      OrderStack
}

func NewRandomizer() *Randomizer {
	pool := Randomizer{}

	pool.Tetrominos = []Tetromino{
		I_PIECE, J_PIECE, L_PIECE, O_PIECE, S_PIECE, T_PIECE, Z_PIECE,
		I_PIECE, J_PIECE, L_PIECE, O_PIECE, S_PIECE, T_PIECE, Z_PIECE,
		I_PIECE, J_PIECE, L_PIECE, O_PIECE, S_PIECE, T_PIECE, Z_PIECE,
		I_PIECE, J_PIECE, L_PIECE, O_PIECE, S_PIECE, T_PIECE, Z_PIECE,
		I_PIECE, J_PIECE, L_PIECE, O_PIECE, S_PIECE, T_PIECE, Z_PIECE,
	}

	pool.history = HistoryStack{S_PIECE, Z_PIECE, S_PIECE}

	return &pool
}

// func (p TGM3Pool)

func (p *Randomizer) Next() Tetromino {
	if len(p.order) == 0 {
		firstPiece := []Tetromino{I_PIECE, J_PIECE, L_PIECE, T_PIECE}[rand.Intn(4)]
		p.history.Push(firstPiece)
		return firstPiece
	} else {
		var piece Tetromino
		var index int
		// Roll for piece
		for roll := 0; roll < 6; roll++ {
			index = rand.Intn(len(p.Tetrominos))
			piece = p.Tetrominos[index]
			if !p.history.Contains(piece) || roll == 5 {
				break
			}
			if len(p.order) > 0 {
				p.Tetrominos[index] = p.order.Peek()
			}
		}

		p.order.Push(piece)

		p.Tetrominos[index] = p.order.Peek()

		// Update history
		p.history.Push(piece)

		return piece
	}
}

//   let pieces = ['I', 'J', 'L', 'O', 'S', 'T', 'Z'];
//   let order = [];

// Create 35 pool.
//   let pool = pieces.concat(pieces, pieces, pieces, pieces);

// First piece special conditions
//   const firstPiece = ['I', 'J', 'L', 'T'][Math.floor(Math.random() * 4)];
//   yield firstPiece;

//   let history = ['S', 'Z', 'S', firstPiece];

//   while (true) {
//     let roll;
//     let i;
//     let piece;

// Roll For piece
//     for (roll = 0; roll < 6; ++roll) {
//       i = Math.floor(Math.random() * 35);
//       piece = pool[i];
//       if (history.includes(piece) === false || roll === 5) {
//         break;
//       }
//       if (order.length) pool[i] = order[0];
//     }

// Update piece order
//     if (order.includes(piece)) {
//       order.splice(order.indexOf(piece), 1);
//     }
//     order.push(piece);

//     pool[i] = order[0];

// Update history
//     history.shift();
//     history[3] = piece;

//     yield piece;
//   }
//}
