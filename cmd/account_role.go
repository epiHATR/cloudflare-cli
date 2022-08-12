/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/consts/text"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// accountRoleCmd represents the accountRole command
var accountRoleCmd = &cobra.Command{
	Use:   "role",
	Short: "manage Cloudflare account/organization roles",
	Long:  text.AccountRoleLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, text.EmptyArgsText+text.SubCmdHelpText)
			os.Exit(1)
		}
		fmt.Print(cmd.Aliases)
	},
}

func init() {
	accountCmd.AddCommand(accountRoleCmd)
}
