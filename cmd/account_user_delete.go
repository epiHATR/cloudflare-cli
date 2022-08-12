/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/account"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var accountUserDeleteAccountId string = ""
var accountUserDeleteUserId string = ""
var userForceDelete bool = false
var userDeleteConfirmText string = "no"

// userDeleteCmdCmd represents the userDeleteCmd command
var userDeleteCmdCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete an user from a Cloudflare account/organization",
	Long:  text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if accountUserDeleteAccountId == "" {
			errText = append(errText, "--account-id")
		}
		if accountUserDeleteUserId == "" {
			errText = append(errText, "--user-id")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		if !userForceDelete {
			for {
				fmt.Print("Please confirm your action (yes/no): ")
				fmt.Scan(&userDeleteConfirmText)
				if strings.ToLower(userDeleteConfirmText) == "yes" || strings.ToLower(userDeleteConfirmText) == "no" {
					break
				}
			}
		} else {
			userDeleteConfirmText = "yes"
		}

		if userDeleteConfirmText == "yes" {
			res := account.AccountDeleteUser(accountUserDeleteAccountId, accountUserDeleteUserId)
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to delete user. The error is", res.Errors[0].Message)
				os.Exit(1)
			} else {
				result := response.CmdResponse{}
				result.Success = res.Success
				output.PrintOut(result, flagQuery, flagOutput)
			}
		}
	},
}

func init() {
	accountUserCmd.AddCommand(userDeleteCmdCmd)
	userDeleteCmdCmd.Flags().StringVarP(&accountUserDeleteAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
	userDeleteCmdCmd.Flags().StringVarP(&accountUserDeleteUserId, "user-id", "", "", "ID of Cloudflare account/organization")
	userDeleteCmdCmd.Flags().BoolVarP(&userForceDelete, "--force", "f", false, "force delete user without confirm")
}
