package main

import (
	"github.com/Chanadu/better-type/btype"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		// btype.NewSimplePage("Hello, World!"),
		btype.CreateTestInput("testing", "teinst "),
	)

	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
