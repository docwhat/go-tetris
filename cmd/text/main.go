package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	foreground  = lipgloss.AdaptiveColor{Light: "#383838", Dark: "#D9DCCF"}
	titleBorder = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	width       = 40
	height      = 23
	titleStyle  = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Margin(0, 2, 0, 0).
			Foreground(foreground).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(titleBorder)
	itemStyle = lipgloss.NewStyle().
			Margin(0, 2).
			Foreground(foreground)
	selectedItemStyle = lipgloss.NewStyle().
				Margin(0, 2).
				Foreground(foreground)
	quitStyle = lipgloss.NewStyle().
			Foreground(foreground)

	docStyle = lipgloss.NewStyle().
			Width(width).
			Height(height).
			BorderStyle(lipgloss.NormalBorder()).
			Foreground(foreground)
)

type model struct {
	list    list.Model
	playing gameMode
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := strings.ToLower(msg.String()); keypress {
		case "ctrl+c", "esc", "q":
			log.Printf("quit: %v", keypress)
			m.playing = QuitGame
			return m, tea.Quit
		case "enter":
			item := m.list.SelectedItem()
			m.playing = item.(gameMode)
		default:
			log.Printf("keypress: %v", keypress)
		}
	case tea.WindowSizeMsg:
		m.list.Styles.Title.Width(width - docStyle.GetHorizontalFrameSize()*2 - m.list.Styles.Title.GetHorizontalMargins())
		// 	x, _ := docStyle.GetFrameSize()
		// 	x2, _ := m.list.Styles.Title.GetFrameSize()
		// 	m.list.Styles.Title.Width(width - x - x2*2)
		// 	// 	m.list.SetSize(msg.Width-x, msg.Height-y)
	}

	nextList, cmd := m.list.Update(msg)
	m.list = nextList
	return m, cmd
}

func (m model) View() string {
	if m.playing == QuitGame {
		return quitStyle.Render("Прощай!")
	}

	// return docStyle.Render(m.list.View())
	return docStyle.Render(m.list.View())
}

type gameMode uint

const (
	ClassicGame gameMode = iota
	MarathonGame
	QuitGame
)

func (m gameMode) String() string {
	switch m {
	case ClassicGame:
		return "Classic Tetris"
	case MarathonGame:
		return "Infinite Tetris"
	case QuitGame:
		return "Quit"
	}
	return ""
}

func (m gameMode) Title() string       { return m.String() }
func (m gameMode) Description() string { return "" }
func (m gameMode) FilterValue() string { return "" }

func NewMenuModel() tea.Model {
	items := []list.Item{
		ClassicGame,
		MarathonGame,
		QuitGame,
	}

	delegate := list.NewDefaultDelegate()
	delegate.SetSpacing(0)
	delegate.ShowDescription = false

	l := list.New(items, delegate, 0, height)
	l.Title = "TETRIS"
	l.Styles.Title = titleStyle
	l.SetFilteringEnabled(false)
	l.SetShowStatusBar(false)
	l.SetShowPagination(false)

	return model{list: l}
}

func menuLoop() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := tea.NewProgram(NewMenuModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
