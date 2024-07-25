
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
  "CVGolang/models"
)

func main() {
	// Initialize the program
	p := tea.NewProgram(models.InitialModel(), tea.WithAltScreen())

	// Run the program
	if err := p.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

