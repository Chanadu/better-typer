package btype

import (
	"fmt"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	text      string
	err       error
}

func CreateTestInput(text, placeholder string) model {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.ShowSuggestions = true
	ti.Focus()
	ti.CharLimit = 0
	ti.Width = 0

	return model{
		textInput: ti,
		text:      text,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	// return textinput.Blink
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput.Cursor.SetMode(cursor.CursorHide)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	default:
		m.textInput.Width += 1

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	view := m.textInput.View()
	placementText := view

	// if m.textInput.Cursor.View() == "" {
	// 	placementText += " "
	// }
	//
	// placementText += m.text

	return fmt.Sprintf(
		"What’s your favorite Pokémon?\n\n%s\n\n",
		placementText,
	) + "\n"
}
