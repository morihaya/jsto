/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// istCmd represents the ist command
var istCmd = &cobra.Command{
	Use:   "ist",
	Short: "show IST time (UTC+5:30, JST-3:30)",
	Long: `Displays the time in edt. This is -3:30 hours from Japan time.

ex)
'IST' time is:
 2022/04/22 13:12:45
	`,
	Run: func(cmd *cobra.Command, args []string) {
		loc, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			panic(err)
		}
		t := time.Now().In(loc)
		fmt.Println("'IST' time is (UTC+5:30, JST-3:30):\n", t.Format("2006/01/02 15:04:05"))
	},
}

func init() {
	rootCmd.AddCommand(istCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// istCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// istCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
