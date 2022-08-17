/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/cache"
	"cloudflare/pkg/consts/jsonBody"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var purgeCmdZoneId string = ""
var purgeCmdPurgeEverything bool = false

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge caches on Cloudflare zone",
	Long:  text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(purgeCmdZoneId) <= 0 {
			errText = append(errText, "--zone-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		if purgeCmdPurgeEverything {
			// call purge everything on Cloudflare zone
			res := cache.PurgeEverything(purgeCmdZoneId, jsonBody.PurgeEverythingJson)
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to purge everything on zone. The error is", res.Errors[0].Message)
				os.Exit(1)
			}
			result := response.CmdResponse{}
			result.Success = res.Success
			output.PrintOut(result, flagQuery, flagOutput)
		}
	},
}

func init() {
	cacheCmd.AddCommand(purgeCmd)
	purgeCmd.Flags().StringVarP(&purgeCmdZoneId, "zone-id", "", "", "ID of cloudfflare zone neeed to purge")
	purgeCmd.Flags().BoolVarP(&purgeCmdPurgeEverything, "purge-everything", "A", false, "Purge everything on Cloudflare zone")
}
