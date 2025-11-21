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

// utcCmd represents the utc command
var utcCmd = &cobra.Command{
	Use:   "utc",
	Short: "show UTC time (UTC+0, JST-9)",
	Long: `Displays the time in UTC. This is -9 hours from Japan time.

ex)
'UTC' time is:
 2022/04/22 13:12:45
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var s string
		var err error
		if len(args) > 0 {
			s, err = converter.ConvertFromJST(args[0], "UTC")
		} else {
			t := time.Now()
			s, err = converter.Convert(t, "UTC")
		}
		if err != nil {
			panic(err)
		}
		fmt.Println("'UTC' The time is:\n", s)
	},
}

func init() {
	rootCmd.AddCommand(utcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// utcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// utcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
