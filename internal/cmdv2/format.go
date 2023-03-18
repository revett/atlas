package cmdv2

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/samber/lo"
)

// https://www.ditig.com/256-colors-cheat-sheet
type colorCodeANSI256 string

const (
	colorCodeBlue colorCodeANSI256 = "5"
	colorCodeRed  colorCodeANSI256 = "9"
	newLine                        = ""
)

func outputError(err error) string {
	blue := colorStyle(colorCodeBlue)
	red := colorStyle(colorCodeRed)

	title := fmt.Sprintf(`Error: "%s"`, err.Error())
	title = red.SetString(title).String()

	stackTrace := lo.Drop(
		strings.Split(fmt.Sprintf("%+v", err), "\n"), 1,
	)

	lines := []string{
		title,
		newLine,
	}
	lines = append(lines, stackTrace...)
	lines = append(
		lines,
		newLine,
		blue.SetString("Need some help?").String(),
		newLine,
		"• 📘 Read the FAQ: https://github.com/revett/atlas#faq",
		"• 🙋 Create a GitHub issue: https://github.com/revett/atlas/issues",
	)

	// 🐛 Needed as the output to bubbletea.Model.View must end with a newline.
	lines = append(lines, newLine)

	return strings.Join(lines, "\n")
}

func colorStyle(code colorCodeANSI256) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(code))
}
