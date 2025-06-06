package btype

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type simplePage struct {
	text string
}

func NewSimplePage(text string) *simplePage {
	return &simplePage{text: text}
}

func (s simplePage) Init() tea.Cmd {
	return nil
}

func (s simplePage) View() string {
	textLen := len(s.text)

	topAndBottomBar := strings.Repeat("*", textLen+4)

	return fmt.Sprintf(
		"%s\n* %s *\n%s",
		topAndBottomBar, s.text, topAndBottomBar,
	)
}

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		}
	}

	return s, nil
}
