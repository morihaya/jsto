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
		var s string
		var err error
		if len(args) > 0 {
			s, err = converter.ConvertFromJST(args[0], "Asia/Kolkata")
		} else {
			t := time.Now()
			s, err = converter.Convert(t, "Asia/Kolkata")
		}
		if err != nil {
			panic(err)
		}
		if len(args) > 0 {
			fmt.Printf("'IST' time for JST %s is:\n %s\n", args[0], s)
		} else {
			fmt.Println("'IST' time is (UTC+5:30, JST-3:30):\n", s)
		}
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
