package utils

import (
	"strings"
  "fmt"

  "CVGolang/styles"
  "CVGolang/assets"
  "github.com/charmbracelet/lipgloss"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func RenderCard(title, company, timeframe, location, details string) string {
	return styles.Card.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			styles.CardTitle.Render(title),
			styles.CardSubtitle.Render(company),
			styles.CardTimeFrame.Render(fmt.Sprintf("%s %s %s", timeframe, styles.Divider, location)),
			styles.CardDetails.Render(details),
		),
	)
}


func RenderSocial(logo, text string) string {
    logoLines := strings.Split(logo, "\n")
    logoHeight := len(logoLines)
    textLines := strings.Split(text, "\n")
    textHeight := len(textLines)

    // Calculate the padding needed to vertically center the text
    paddingTop := (logoHeight - textHeight) / 2

    paddedText := strings.Repeat("\n", paddingTop) + text
    logoPart := styles.LogoStyle.Render(logo)
    textPart := styles.TextStyle.Render(paddedText)
    return lipgloss.JoinHorizontal(lipgloss.Top, logoPart, textPart)
}



func RenderSkill(name string, proficiency int) string {
    filled := lipgloss.NewStyle().
      Foreground(lipgloss.Color("63")).
      Render(strings.Repeat(assets.FilledCircle, proficiency))

    empty := lipgloss.NewStyle().
      Foreground(lipgloss.Color("240")).
      Render(strings.Repeat(assets.EmptyCircle, assets.MaxProficiency-proficiency))

    return lipgloss.JoinHorizontal(
      lipgloss.Top,
      styles.SkillNameStyle.Render(name),
      styles.ProficiencyStyle.Render(filled+empty),
    )
}

func RenderSkillsTable() string {
    var skillRows []string
    for _, skill := range assets.Skills {
        skillRows = append(skillRows, RenderSkill(skill.Name, skill.Proficiency))
    }
    // Add an empty row at the end
    emptyRow := lipgloss.JoinHorizontal(
      lipgloss.Top,
      styles.SkillNameStyle.Render(""),
      styles.ProficiencyStyle.Render(""),
    )

    skillRows = append(skillRows, emptyRow)

    header := lipgloss.JoinHorizontal(
      lipgloss.Left,
      styles.SkillHeaderStyle.Render("Programming Language"),
      styles.ProficiencyHeaderStyle.Render("Proficiency"),
    )

    body := lipgloss.JoinVertical(lipgloss.Left, skillRows...)
    return lipgloss.JoinVertical(lipgloss.Left, header, body)
}


func RenderFrameworksTree() string {
    treeStyle := lipgloss.NewStyle().
      Margin(1, 2).
      Align(lipgloss.Left)

    return treeStyle.Render(assets.FrameworksTree)
}

func RenderSkillsAndFrameworks() string {
    skillsSection := styles.SkillsSectionStyle.Render(assets.LanguageArt + "\n\n" + RenderSkillsTable())
    frameworksSection := styles.FrameworksSectionStyle.Render(assets.FrameworkArt + "\n" + RenderFrameworksTree())
    extraSection := styles.ExtraSectionStyle.Render(assets.ExtraArt + "\n\n" + RenderExtraSection())

    leftDivider := RenderVerticalDivider(19)
    midLeftDivider := RenderVerticalDivider(31)
    midRightDivider := RenderVerticalDivider(31)
    rightDivider := RenderVerticalDivider(11)

    combined := lipgloss.JoinHorizontal(
        lipgloss.Top,
        leftDivider,
        skillsSection,
        midLeftDivider,
        frameworksSection,
        midRightDivider,
        extraSection,
        rightDivider,
    )
    return combined
}


func RenderVerticalDivider(height int) string {
    verticalLine := ""
    for i := 0; i < height; i++ {
        verticalLine += "│\n"
    }
    return verticalLine
}


func RenderExtraSection() string {
    // Define the data
    languages := []struct {
        Language   string
        Proficiency string
    }{
        {"English", "Fluent"},
        {"Serbian", "Fluent"},
        {"German", "Intermediate"},
    }

    // Create the formatted strings
    var formattedText string
    formattedText += styles.ExtraHeaderStyle.Render("Languages") + "\n"
    for _, lang := range languages {
        formattedText += styles.ExtraTextStyle.Render(fmt.Sprintf("%s: %s", lang.Language, lang.Proficiency)) + "\n"
    }

    return formattedText
}

