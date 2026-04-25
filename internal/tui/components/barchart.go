package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BarItem represents a single bar in a bar chart.
type BarItem struct {
	Label string
	Value int
}

// BarChart renders a horizontal bar chart with colored bars.
func BarChart(items []BarItem, maxWidth int) string {
	if len(items) == 0 {
		return ""
	}

	maxLabel := 0
	maxVal := 0
	for _, item := range items {
		if len(item.Label) > maxLabel {
			maxLabel = len(item.Label)
		}
		if item.Value > maxVal {
			maxVal = item.Value
		}
	}

	barWidth := maxWidth - maxLabel - 10
	if barWidth > 40 {
		barWidth = 40
	}
	if barWidth < 10 {
		barWidth = 10
	}

	labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF"))
	countStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF")).Faint(true)

	var lines []string
	for _, item := range items {
		// Use half-block resolution: each character position = 2 units
		var units int
		if maxVal > 0 {
			units = item.Value * barWidth * 2 / maxVal
		}
		if units < 1 && item.Value > 0 {
			units = 1
		}
		fullBlocks := units / 2
		halfBlock := units % 2

		// Color based on value relative to maxVal
		var color string
		ratio := 0.0
		if maxVal > 0 {
			ratio = float64(item.Value) / float64(maxVal)
		}
		switch {
		case ratio >= 0.75:
			color = "#FED90F"
		case ratio >= 0.50:
			color = "#FF8C00"
		case ratio >= 0.25:
			color = "#F59E0B"
		default:
			color = "#B85C00"
		}
		barStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(color))

		barStr := strings.Repeat("█", fullBlocks)
		if halfBlock == 1 {
			barStr += "▌"
		}

		label := labelStyle.Render(fmt.Sprintf("%-*s", maxLabel, item.Label))
		coloredBar := barStyle.Render(barStr)
		count := countStyle.Render(fmt.Sprintf("%d", item.Value))
		line := fmt.Sprintf("%s %s %s", label, coloredBar, count)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}
