/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/yimikao/accountant/database"
)

//variables to store flag values
var CreditAmount float64
var CreditDetails string

// creditCmd represents the credit command
var creditCmd = &cobra.Command{
	Use:   "credit",
	Short: "Create a credit transaction.",
	Long: `This command allows to create a debit transaction for a particular user.
Usage: accountant credit <username> --amount=<amount> --details:<details>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Username not speccified.")
		}

		username := args[0]

		u, err := database.FindUser(username)
		if err != nil {
			log.Fatal(err)
		}
		u.UpdateAmount(CreditAmount, CreditDetails)

	},
}

func init() {
	rootCmd.AddCommand(creditCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// creditCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// creditCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	creditCmd.Flags().StringVarP(&CreditDetails, "details", "d", "", "details of what transaction is about")
	creditCmd.Flags().Float64VarP(&CreditAmount, "amount", "a", 0, "amount credited for user.")

	creditCmd.MarkFlagRequired("amount")
	creditCmd.MarkFlagRequired("details")

}
