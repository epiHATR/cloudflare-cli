/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/account"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"

	"github.com/spf13/cobra"
)

var accountCmdAccountName string = ""

// listCmd represents the list command
var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all managed Cloudflare accounts/organizations",
	Long:  text.AccountListLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			result := []response.AccountDetails{}
			response := account.GetAllAccounts(1, accountCmdAccountName)
			result = append(result, response.Result...)
			output.PrintOut(result, flagQuery, flagOutput)
		}
	},
}

func init() {
	accountCmd.AddCommand(accountListCmd)
	accountListCmd.Flags().StringVarP(&accountCmdAccountName, "name", "n", "", "name of account to search")
}
