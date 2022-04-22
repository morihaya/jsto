/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// estCmd represents the edt command
var estCmd = &cobra.Command{
	Use:   "edt",
	Short: "show EDT time (UTC-4, JST-13)",
	Long: `Displays the time in edt. This is -13 hours from Japan time.

ex)
'EDT' time is:
 2022/04/22 13:12:45
	`,
	Run: func(cmd *cobra.Command, args []string) {
		loc, err := time.LoadLocation("America/New_York")
		if err != nil {
			panic(err)
		}
		t := time.Now().In(loc)
		fmt.Println("'EDT' time is (UTC-4, JST-13):\n", t.Format("2006/01/02 15:04:05"))
	},
}

func init() {
	rootCmd.AddCommand(estCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// estCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// estCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
