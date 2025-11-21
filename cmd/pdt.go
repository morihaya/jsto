/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/morihaya/jsto/pkg/converter"
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
		var s string
		var err error
		if len(args) > 0 {
			s, err = converter.ConvertFromJST(args[0], "America/Los_Angeles")
		} else {
			t := time.Now()
			s, err = converter.Convert(t, "America/Los_Angeles")
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("'PDT' The time is:\n", s)
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
