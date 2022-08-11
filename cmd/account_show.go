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

var accountShowCmdAccountId string = ""

// accountShowCmd represents the show command
var accountShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show Cloudflare account/organization details",
	Long:  text.AccountShowLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(accountShowCmdAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := account.GetAccountDetails(accountShowCmdAccountId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get account/organization details. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountCmd.AddCommand(accountShowCmd)
	accountShowCmd.Flags().StringVarP(&accountShowCmdAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
}
