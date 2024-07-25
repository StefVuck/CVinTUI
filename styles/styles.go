package styles

import "github.com/charmbracelet/lipgloss"

/* 
Constants 
----------
These constants define fixed values used throughout the styling.
*/
const (
	ColumnWidth    = 60
	MaxProficiency = 10
)

/* 
Adaptive Colors 
---------------
These are adaptive colors used for styling elements depending on light or dark mode.
*/
var (
	Subtle    = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#66b2b2"}
	Highlight = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#FF69B4"}
	Special   = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
)

/* 
Basic Elements
--------------
These variables define basic styled elements using the lipgloss package.
*/
var (
	Divider = lipgloss.NewStyle().
		SetString("•").
		Padding(0, 1).
		Foreground(Subtle).
		String()

	URL = lipgloss.NewStyle().
		Foreground(Special).
		Render

	ActiveTab = lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─", Bottom: " ", Left: "│", Right: "│", TopLeft: "╭", TopRight: "╮", BottomLeft: "┘", BottomRight: "└"}, true).
		BorderForeground(Highlight).
		Padding(0, 1)

	Tab = lipgloss.NewStyle().
		Border(lipgloss.Border{Top: "─", Bottom: "─", Left: "│", Right: "│", TopLeft: "╭", TopRight: "╮", BottomLeft: "┴", BottomRight: "┴"}, true).
		BorderForeground(Highlight).
		Padding(0, 1)

	TabGap = Tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

/* 
Content Styles 
--------------
These variables define styles for different content elements.
*/
var (
	ContentStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")).
		Padding(1, 2).
		Margin(1, 0).
		Align(lipgloss.Center)

	FooterStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")).
		Align(lipgloss.Center).
		Margin(1, 0)

	NavbarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF69B4")).
		Align(lipgloss.Center).
		PaddingBottom(1)

	LogoStyle = lipgloss.NewStyle().
		Width(20).
		Height(12).
		Align(lipgloss.Left)

	TextStyle = lipgloss.NewStyle().
		Align(lipgloss.Left).
		PaddingLeft(2)

	SocialStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		Padding(0, 1).
		Margin(1, 0)
)

/* 
Section Styles 
--------------
These variables define styles for different sections.
*/
var (
	SkillsSectionStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Left).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderBottom(true)

	FrameworksSectionStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Left).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderBottom(true)

	ExtraSectionStyle = lipgloss.NewStyle().
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Left).
		BorderStyle(lipgloss.NormalBorder()).
		BorderTop(true).
		BorderBottom(true)
)

/* 
Skill Styles 
------------
These variables define styles for skill-related elements.
*/
var (
	SkillNameStyle = lipgloss.NewStyle().
		Width(20).
		Align(lipgloss.Left).
		PaddingRight(4)

	ProficiencyStyle = lipgloss.NewStyle().
		Align(lipgloss.Left).
		PaddingLeft(2)

	SkillHeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Width(20).
		Align(lipgloss.Left).
		PaddingRight(4)

	ProficiencyHeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Align(lipgloss.Left).
		PaddingLeft(2)

	ExtraHeaderStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))

	ExtraTextStyle = lipgloss.NewStyle().
		PaddingLeft(2)
)

/* 
Tree Styles 
-----------
These variables define styles for tree-related elements.
*/
var (
	TreeStyle = lipgloss.NewStyle().
		Margin(1, 2)

	RootStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))
)

/* 
Card Styles 
-----------
These variables define styles for card-related elements.
*/
var (
	CardTitle     = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFF")).
		PaddingBottom(1)

	CardSubtitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#C4C8E2")).
		PaddingBottom(1)

	CardTimeFrame = lipgloss.NewStyle().
		Foreground(Subtle).
		PaddingBottom(1)

	CardDetails = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF")). 
		PaddingTop(1)

	Card = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#66b2b2")). 
		Padding(1).
		MarginBottom(1).
		Width(ColumnWidth)
)

