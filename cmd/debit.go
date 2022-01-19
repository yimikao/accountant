/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yimikao/accountant/database"
)

//vars to store debit flags
var DebitAmount float64
var DebitDetails string

func checkBalance(amount float64, balance float64) string {
	if balance-amount < 0 {
		return "insufficient balance"
	}
	return ""
}

// debitCmd represents the debit command
var debitCmd = &cobra.Command{
	Use:   "debit",
	Short: "Create a debit transaction.",
	Long: `This command allows to create a debit transaction for a particular user.
Usage: accountant debit <username> --amount=<amount> --details:<details>`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			log.Fatal("Username not specified.")
		}

		username := args[0]

		u, err := database.FindUser(username)

		if err != nil {
			log.Fatal(err)
		}
		status := checkBalance(DebitAmount, u.Balance)
		if len(status) > 0 {
			log.Fatal(status)
			return
		}
		u.Update(-DebitAmount, DebitDetails)
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

	debitCmd.Flags().Float64VarP(&DebitAmount, "amount", "a", 0, "amount to deduct from user's balance")
	debitCmd.Flags().StringVarP(&DebitDetails, "details", "d", "", "what the amount was deducted for.")

	debitCmd.MarkFlagRequired("amount")
	debitCmd.MarkFlagRequired("details")
}
