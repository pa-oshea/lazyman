package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

const listHeight = 30

type itemDelegate struct {
	styles PackagesStyle
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Package)
	if !ok {
		return
	}

	fmt.Fprint(w, " ")
	if index == m.Index() {
		fmt.Fprint(w, d.styles.Selected.Render("> "+i.String()))
		return
	}

	fmt.Fprint(w, d.styles.Unselected.Render("  "+i.String()))
}

type Model struct {
	// The array of packages that will be listed
	Packages list.Model
	// Style for list
	PackagesStyle PackagesStyle
	// Style for the package information
	informationStyle lipgloss.Style
	helpStyle        lipgloss.Style
	// The package we are currently hightlighting in the list
	choice Package
	// informationView
	informationView viewport.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.Packages.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.Packages, cmd = m.Packages.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			m.PackagesStyle.Base.Render(m.Packages.View()),
			m.informationStyle.Border(lipgloss.DoubleBorder()).Render(m.informationView.View())),
	)
}
