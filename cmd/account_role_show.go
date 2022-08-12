/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/account"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var accountRoleShowAccountId string = ""
var accountRoleShowRoleId string = ""

// accountRoleShowCmd represents the accountRoleShow command
var accountRoleShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details of a Cloudflare role",
	Long:  text.AccountRoleDetailsLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(accountRoleShowAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(accountRoleShowRoleId) <= 0 {
			errText = append(errText, "--role-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		res := account.GetAccountRoleDetails(accountRoleShowAccountId, accountRoleShowRoleId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get details of Cloudflare role. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountRoleCmd.AddCommand(accountRoleShowCmd)
	accountRoleShowCmd.Flags().StringVarP(&accountRoleShowAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
	accountRoleShowCmd.Flags().StringVarP(&accountRoleShowRoleId, "role-id", "", "", "user ID")
}
