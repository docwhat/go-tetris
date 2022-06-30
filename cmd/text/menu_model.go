package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleBorder = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	titleStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Margin(0, 2, 0, 0).
			Foreground(foreground).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(titleBorder)

	quitStyle = lipgloss.NewStyle().
			Width(width).
			Margin(2, 2).
			Align(lipgloss.Center).
			Foreground(foreground)

	docStyle = lipgloss.NewStyle().
			Width(width).
			Height(height).
			BorderStyle(lipgloss.NormalBorder()).
			Foreground(foreground)
)

type menuModel struct {
	list    list.Model
	playing gameMode
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := strings.ToLower(msg.String()); keypress {
		case "ctrl+c", "esc", "q":
			log.Printf("quit: %v", keypress)
			m.playing = QuitGame
			return m, tea.Quit
		case "enter", "space":
			m.playing = m.list.SelectedItem().(gameMode)
			log.Printf("selecting %v", m.playing)
			if m.playing == QuitGame {
				return m, tea.Quit
			}
		default:
			m.playing = Nothing
			log.Printf("keypress: %v", keypress)
		}
	case tea.WindowSizeMsg:
		m.list.Styles.Title.Width(width - docStyle.GetHorizontalFrameSize()*2 - m.list.Styles.Title.GetHorizontalMargins())
	}

	nextList, cmd := m.list.Update(msg)
	m.list = nextList
	return m, cmd
}

func (m menuModel) View() string {
	// log.Printf("selecting %v", m.playing)
	if m.playing == QuitGame {
		return quitStyle.Render("Прощай!\n")
	}
	return docStyle.Render(m.list.View())
}

func NewMenuModel() tea.Model {
	items := []list.Item{
		ClassicGame,
		// MarathonGame,
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

	return menuModel{list: l}
}
