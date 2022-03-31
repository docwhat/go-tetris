package main

type gameMode uint

const (
	ClassicGame gameMode = iota
	MarathonGame
	QuitGame
	Nothing
)

func (m gameMode) String() string {
	switch m {
	case ClassicGame:
		return "Classic Tetris"
	case MarathonGame:
		return "Infinite Tetris"
	case QuitGame:
		return "Quit"
	case Nothing:
		return "Nothing"
	}
	return ""
}

func (m gameMode) Title() string       { return m.String() }
func (m gameMode) Description() string { return "" }
func (m gameMode) FilterValue() string { return "" }
