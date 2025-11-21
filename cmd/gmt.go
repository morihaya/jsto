/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/morihaya/jsto/pkg/converter"
	"github.com/spf13/cobra"
)

// gmtCmd represents the gmt command
var gmtCmd = &cobra.Command{
	Use:   "gmt",
	Short: "show GMT/BST time (UTC+0/+1, JST-9/-8)",
	Long: `Displays the time in GMT/BST (UK Time). This is -9 hours (GMT) or -8 hours (BST) from Japan time.

ex)
'GMT' time is:
 2022/04/22 13:12:45

You can also specify a time in JST to convert:
$ jsto gmt 12:30
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var s string
		var err error
		if len(args) > 0 {
			s, err = converter.ConvertFromJST(args[0], "Europe/London")
		} else {
			t := time.Now()
			s, err = converter.Convert(t, "Europe/London")
		}
		if err != nil {
			panic(err)
		}
		if len(args) > 0 {
			fmt.Printf("'GMT' time for JST %s is:\n %s\n", args[0], s)
		} else {
			fmt.Println("'GMT' time is (UTC+0/+1, JST-9/-8):\n", s)
		}
	},
}

func init() {
	rootCmd.AddCommand(gmtCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gmtCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gmtCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
