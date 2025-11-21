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

// cstCmd represents the cst command
var cstCmd = &cobra.Command{
	Use:   "cst",
	Short: "show CST time (UTC+8, JST-1)",
	Long: `Displays the time in CST (China Standard Time). This is -1 hour from Japan time.

ex)
'CST' time is:
 2022/04/22 13:12:45

You can also specify a time in JST to convert:
$ jsto cst 12:30
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var s string
		var err error
		if len(args) > 0 {
			s, err = converter.ConvertFromJST(args[0], "Asia/Shanghai")
		} else {
			t := time.Now()
			s, err = converter.Convert(t, "Asia/Shanghai")
		}
		if err != nil {
			panic(err)
		}
		if len(args) > 0 {
			fmt.Printf("'CST' time for JST %s is:\n %s\n", args[0], s)
		} else {
			fmt.Println("'CST' time is (UTC+8, JST-1):\n", s)
		}
	},
}

func init() {
	rootCmd.AddCommand(cstCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cstCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cstCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
