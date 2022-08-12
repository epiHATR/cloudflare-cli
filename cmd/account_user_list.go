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

var accountUserListAccountId string = ""

// accountUserListCmd represents the list command
var accountUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all user in a Cloudflare account/organization",
	Long:  text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if accountUserListAccountId == "" {
			errText = append(errText, "--account-id")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := account.GetAccountUsers(accountUserListAccountId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get users of Cloudflare account/organization. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountUserCmd.AddCommand(accountUserListCmd)
	accountUserListCmd.Flags().StringVarP(&accountUserListAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
}
