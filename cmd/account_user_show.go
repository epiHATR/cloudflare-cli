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

var accountUserDetailAccountId string = ""
var accountUserDetailId string = ""

// accountUserDetailCmd represents the show command
var accountUserDetailCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user's information details",
	Long:  text.AccountUserDetailsLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(accountUserDetailAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(accountUserDetailId) <= 0 {
			errText = append(errText, "--user-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := account.GetAccountUserDetail(accountUserDetailAccountId, accountUserDetailId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get user details. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountUserCmd.AddCommand(accountUserDetailCmd)
	accountUserDetailCmd.Flags().StringVarP(&accountUserDetailAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
	accountUserDetailCmd.Flags().StringVarP(&accountUserDetailId, "user-id", "", "", "user ID")
}
