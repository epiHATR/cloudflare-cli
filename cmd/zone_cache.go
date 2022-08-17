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

// cacheCmdCmd represents the cacheCmd command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "control cache behavior on Cloudflare zone",
	Long:  text.CacheCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, text.EmptyArgsText+text.SubCmdHelpText)
			os.Exit(1)
		}
		fmt.Print(cmd.Aliases)
	},
}

func init() {
	zoneCmd.AddCommand(cacheCmd)
}
