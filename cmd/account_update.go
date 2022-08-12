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
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var accountUpdateCmdAccountId string = ""
var accountUpdateData string = ""
var accountForceUpdate bool = false
var accountUpdateConfirmText = "no"

var accountUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a Cloudflare account/organization information",
	Long:  text.AccountUpdateLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(accountUpdateCmdAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}
		if len(accountUpdateData) <= 0 {
			errText = append(errText, "--data|-d")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		log.Println("Updating account id", accountUpdateCmdAccountId)
		if !accountForceUpdate {
			for {
				fmt.Print("Please confirm your action (yes/no): ")
				fmt.Scan(&accountUpdateConfirmText)
				if strings.ToLower(accountUpdateConfirmText) == "yes" || strings.ToLower(accountUpdateConfirmText) == "no" {
					break
				}
			}
		} else {
			accountUpdateConfirmText = "yes"
		}

		if accountUpdateConfirmText == "yes" {
			res := account.UpdateAccount(accountUpdateCmdAccountId, accountUpdateData)

			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to update account/organization. The error is", res.Errors[0].Message)
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
	accountCmd.AddCommand(accountUpdateCmd)
	accountUpdateCmd.Flags().StringVarP(&accountUpdateCmdAccountId, "account-id", "", "", "ID of Cloudflare account/organization")
	accountUpdateCmd.Flags().StringVarP(&accountUpdateData, "data", "d", "", "DNS id need to update")
	accountUpdateCmd.Flags().BoolVarP(&accountForceUpdate, "--force", "f", false, "force update DNS records without confirm")
}
