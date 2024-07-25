package views

import (
  tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
  "CVGolang/styles"  
  "CVGolang/utils"  
  "strings"
)

type Model struct {
	Index  int
	Pages  []string
	Width  int
	Height int
	Offset int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "left", "h":
			if m.Index > 0 {
				m.Index--
				m.Offset = 0 // Reset offset when changing page
			}
		case "right", "l":
			if m.Index < len(m.Pages)-1 {
				m.Index++
				m.Offset = 0 // Reset offset when changing page
			}
		case "up", "k":
			if m.Offset > 0 {
				m.Offset--
			}
		case "down", "j":
			m.Offset++
		}
	}
	return m, nil
}





func (m Model) View() string {
	// Create the navigation bar
	var navItems []string
	tabs := []string{"Home", "Experience", "Skills", "Projects", "Hobbies", "Socials", "Help"}
	for i := range tabs {
		if i == m.Index {
			navItems = append(navItems, styles.ActiveTab.Render(tabs[i]))
		} else {
			navItems = append(navItems, styles.Tab.Render(tabs[i]))
		}
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, navItems...)
	gap := styles.TabGap.Render(strings.Repeat(" ", utils.Max(0, m.Width-lipgloss.Width(row)-2)))
	navbar := lipgloss.JoinHorizontal(lipgloss.Bottom, row, gap)
	fullNavbar := styles.NavbarStyle.Width(m.Width).Render(navbar)

	// Parse the current page content
	contentParts := m.Pages[m.Index]
	contentWrapped := wordwrap.String(contentParts, m.Width-4)
	contentLines := strings.Split(contentWrapped, "\n")

	// Calculate content height
	contentHeight := m.Height - lipgloss.Height(fullNavbar) - lipgloss.Height(styles.FooterStyle.String()) - 2
	if contentHeight < 0 {
		contentHeight = 0
	}
	if m.Offset > len(contentLines)-contentHeight {
		m.Offset = len(contentLines) - contentHeight
	}
	if m.Offset < 0 {
		m.Offset = 0
	}
	visibleContent := contentLines[m.Offset:utils.Min(m.Offset+contentHeight, len(contentLines))]

	// Create the visible content
	scrollableContent := strings.Join(visibleContent, "\n")

	// Horizontally center the scrollable content
	scrollableContent = lipgloss.PlaceHorizontal(m.Width, lipgloss.Center, scrollableContent)

	// Create the footer
	footer := styles.FooterStyle.Width(m.Width).Render("Press 'q' to quit. Use left/right arrow keys to navigate. Use up/down arrow keys to scroll.")

	// Combine the navigation bar, main content, and footer
	mainLayout := lipgloss.JoinVertical(
		lipgloss.Top,
		fullNavbar,
		scrollableContent,
		footer,
	)

	return mainLayout
}
