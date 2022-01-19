/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// debitCmd represents the debit command
var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "Create a debit transaction.",
	Long: `This command allows to create a debit transaction for a particular user.
Usage: accountant debit <username> --amount=<amount> --details:<details>`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debit called")
	},
}

func init() {
	rootCmd.AddCommand(debitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// debitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// debitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
