package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var urlFlag string

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "you can pass your testing endpoint url with -e and request body file with -b ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(urlFlag)
	},
}

func init() {
	testCmd.Flags().StringVarP(&urlFlag, "address", "a", "", "URL to fuzztest")
	testCmd.MarkFlagRequired("address")
	rootCmd.AddCommand(testCmd)
}
