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

var accountUserAddAccountId string = ""
var accountUserAddEmail string = ""
var accountUserAddStatus string = ""
var accountUserAddRoles []string = nil

// accountUserAddCmd represents the add command
var accountUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add new user to a Cloudflare account/organization",
	Long:  text.AccountUserAddLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(accountUserAddAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(accountUserAddEmail) <= 0 {
			errText = append(errText, "--email")
		}

		if len(accountUserAddStatus) <= 0 {
			errText = append(errText, "--status")
		} else if strings.ToLower(accountUserAddStatus) != "accepted" && strings.ToLower(accountUserAddStatus) != "pending" {
			errText = append(errText, "--status")
		}

		if len(accountUserAddRoles) <= 0 {
			errText = append(errText, "--role-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := account.AccountAddUser(accountUserAddAccountId, accountUserAddEmail, accountUserAddStatus, accountUserAddRoles)
		if !res.Success {
			//we're going to remove failed zone here
			fmt.Fprintln(os.Stderr, "Error: Failed to add user. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	accountUserCmd.AddCommand(accountUserAddCmd)
	accountUserAddCmd.Flags().SortFlags = false
	accountUserAddCmd.Flags().StringVarP(&accountUserAddAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
	accountUserAddCmd.Flags().StringVarP(&accountUserAddEmail, "email", "e", "", "User email need to be added")
	accountUserAddCmd.Flags().StringVarP(&accountUserAddStatus, "status", "", "", "Status when added to a organization")
	accountUserAddCmd.Flags().StringArrayVarP(&accountUserAddRoles, "role-id", "", []string{""}, "Status when added to a organization")
}
