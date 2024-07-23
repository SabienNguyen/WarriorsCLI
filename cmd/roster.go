/*
Copyright © 2024 Sabien Nguyen Sabiennguyen@gmail.com
*/
package cmd

import (
	"fmt"

	"warriors-cli/ui"

	wapi "github.com/SabienNguyen/WAPI"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// rosterCmd represents the roster command
var rosterCmd = &cobra.Command{
	Use:   "roster",
	Short: "Shows the current Warriors player roster",
	Long: `This command shows the current Warriors player roster. 
	When the roster is shown, the user can move up and 
	down the list using the arrow keys. Various information
	about the player is shown on the table. `,
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
			BorderStyle(lipgloss.RoundedBorder()).
			BorderBottomForeground(lipgloss.Color("#ffba00")).
			BorderBottom(true).Bold(false)
		s.Selected = s.Selected.
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("#00408d")).
			Bold(false)
		t.SetStyles(s)

		m := ui.Model{Table: t}
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
