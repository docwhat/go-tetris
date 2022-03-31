package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	foreground = lipgloss.AdaptiveColor{Light: "#383838", Dark: "#D9DCCF"}
	width      = 40
	height     = 23
)

func menuLoop() {
	rand.Seed(time.Now().UTC().UnixNano())

	if err := tea.NewProgram(NewMenuModel()).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
