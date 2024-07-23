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

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Shows the season schedule for the Warriors basketball team",
	Long: `
This command shows the schedule for the Warriors 
basketball team. Shows various information about 
each game in the schedule.`,
	Run: func(cmd *cobra.Command, args []string) {
		schedule, err := wapi.GetSchedule()

		if err != nil {
			fmt.Printf("expected schedule got %s", schedule)
		}

		columns := []table.Column{
			{Title: "Opponent", Width: 8},
			{Title: "Date", Width: 6},
			{Title: "Location", Width: 16},
			{Title: "Status", Width: 14},
			{Title: "Notes", Width: 14},
		}

		rows := make([]table.Row, len(schedule))
		for i, p := range schedule {
			rows[i] = table.Row{
				p.Opponent,
				p.Date,
				p.Location,
				p.Status,
				p.Notes,
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
	rootCmd.AddCommand(scheduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
