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

// listCmd represents the list command
var zonePlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "manage Cloudflare zone plans",
	Long:  text.PlancmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, text.EmptyArgsText+text.SubCmdHelpText)
			os.Exit(1)
		}
		fmt.Print(cmd.Aliases)
	},
}

func init() {
	zoneCmd.AddCommand(zonePlanCmd)
}
