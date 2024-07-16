/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tlowerison/nbastats"
)

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
		client := nbastats.NewClient()
		team := "1610612744"
		season := "2020-21"
		data, err := client.Lineups(&nbastats.LineupsFields{TeamID: &team, Season: &season})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(data)
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
