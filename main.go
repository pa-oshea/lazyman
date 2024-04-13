package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	var items []list.Item

	packages := GetUserInstalledPackages()
	for _, v := range packages {
		items = append(items, v)
	}

	defaultStyles := DefaultStyles()

	packageList := list.New(items, itemDelegate{defaultStyles.Packages}, 0, 0)
	packageList.Title = "Installed Packages"
	packageList.SetShowHelp(true)
	packageList.SetFilteringEnabled(true)
	packageList.SetShowStatusBar(false)
	packageList.DisableQuitKeybindings()
	packageList.Styles.NoItems = lipgloss.NewStyle().Margin(0, 2).Foreground(lipgloss.Color(GrayColor))
	packageList.SetStatusBarItemName("package", "packages")

	viewport := viewport.New(80, 0)
	viewport.Width = 150
	viewport.Height = lipgloss.Height(packages[0].Value)
	viewport.SetContent(packages[0].Value)

	m := Model{Packages: packageList, PackagesStyle: defaultStyles.Packages, informationStyle: defaultStyles.Information, informationView: viewport}

	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
