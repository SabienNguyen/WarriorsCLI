/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	wapi "github.com/SabienNguyen/WAPI"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Selected player: %s", m.table.SelectedRow()[0]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

// rosterCmd represents the roster command
var rosterCmd = &cobra.Command{
	Use:   "roster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		roster, err := wapi.GetRoster()

		if err != nil {
			fmt.Printf("expected roster got %s", roster)
		}

		columns := []table.Column{
			{Title: "Name", Width: 20},
			{Title: "Number", Width: 6},
			{Title: "Position", Width: 8},
			{Title: "Height", Width: 8},
			{Title: "Weight", Width: 8},
			{Title: "Birthdate", Width: 12},
			{Title: "Age", Width: 6},
			{Title: "Experience", Width: 12},
			{Title: "School", Width: 16},
			{Title: "Aquired", Width: 32},
		}

		rows := make([]table.Row, len(roster))
		for i, p := range roster {
			rows[i] = table.Row{
				p.Name,
				p.Number,
				p.Position,
				p.Height,
				p.Weight,
				p.Birthdate,
				p.Age,
				p.Experience,
				p.School,
				p.Aquired,
			}
		}

		t := table.New(
			table.WithColumns(columns),
			table.WithRows(rows),
			table.WithFocused(true),
			table.WithHeight(12),
		)

		s := table.DefaultStyles()
		s.Header = s.Header.
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottomForeground(lipgloss.Color("240")).
			BorderBottom(true).Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("57")).
			Bold(false)
		t.SetStyles(s)

		m := model{t}
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fmt.Printf("Error running program: %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rosterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rosterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rosterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
