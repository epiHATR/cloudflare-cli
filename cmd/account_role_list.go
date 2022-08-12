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

var accountRoleListAccountId string = ""

// accountrolelistCmd represents the accountrolelist command
var accountRoleListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all available role in Cloudflare account/organization",
	Long:  text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if accountRoleListAccountId == "" {
			errText = append(errText, "--account-id")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := account.GetAccountRoles(accountRoleListAccountId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get users of Cloudflare account/organization. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountRoleCmd.AddCommand(accountRoleListCmd)
	accountRoleListCmd.Flags().StringVarP(&accountRoleListAccountId, "account-id", "", "", "ID of a Cloudflare account/organization")
}
