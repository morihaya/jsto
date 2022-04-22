/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// pstCmd represents the pst command
var pstCmd = &cobra.Command{
	Use:   "pdt",
	Short: "show PDT time (UTC-7, JST-15)",
	Long: `Displays the time in PDT. This is -15 hours from Japan time.

ex)
'PDT' time is:
 2022/04/22 13:12:45
	`,
	Run: func(cmd *cobra.Command, args []string) {
		loc, err := time.LoadLocation("America/Los_Angeles")
		if err != nil {
			panic(err)
		}
		t := time.Now().In(loc)
		fmt.Println("'PDT' The time is:\n", t.Format("2006/01/02 15:04:05"))
	},
}

func init() {
	rootCmd.AddCommand(pstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
