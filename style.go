package main

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	PrimaryColor        = "#AFBEE1"
	PrimaryColorSubdued = "#64708D"
	BrightGreenColor    = "#BCE1AF"
	GreenColor          = "#527251"
	BrightRedColor      = "#E49393"
	RedColor            = "#A46060"
	ForegroundColor     = "15"
	BackgroundColor     = "235"
	GrayColor           = "241"
	BlackColor          = "#373b41"
	WhiteColor          = "#FFFFFF"
)

// PackagesStyle holds the neccessary styling for the package list
type PackagesStyle struct {
	Base            lipgloss.Style
	Title           lipgloss.Style
	TitleBar        lipgloss.Style
	Selected        lipgloss.Style
	Unselected      lipgloss.Style
	PaginationStyle lipgloss.Style
	QuitTextStyle   lipgloss.Style
}

// Styles is the struct of all styles for the application.
type Styles struct {
	Packages    PackagesStyle
	Information lipgloss.Style
	HelpStyle   lipgloss.Style
}

var marginstyle = lipgloss.NewStyle().Margin(1, 0, 0, 1)

// DefaultStyles is the default implementation of the styles struct for all
// styling in the application.
func DefaultStyles() Styles {
	white := lipgloss.Color(WhiteColor)
	black := lipgloss.Color(BlackColor)
	selected := lipgloss.Color(GreenColor)

	return Styles{
		Packages: PackagesStyle{
			Base:            lipgloss.NewStyle().Margin(1, 2),
			Title:           lipgloss.NewStyle().Padding(0, 1).Foreground(black),
			TitleBar:        lipgloss.NewStyle().Background(selected).Width(22-2).Margin(0, 1, 1, 1),
			Selected:        lipgloss.NewStyle().Foreground(selected),
			Unselected:      lipgloss.NewStyle().Foreground(white),
			PaginationStyle: list.DefaultStyles().PaginationStyle.PaddingLeft(4),
			QuitTextStyle:   lipgloss.NewStyle().Margin(1, 0, 2, 4),
		},
		Information: lipgloss.NewStyle().Margin(0, 1),
		HelpStyle:   list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1),
	}
}
