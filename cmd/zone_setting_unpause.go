/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/setting"
	"cloudflare/pkg/consts/jsonBody"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var unPauseCmdZoneId = ""

// pauseCmd represents the pause command
var unPauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "unpause a Cloudflare zone",
	Long:  text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("verify flag inputs")
		errText := []string{}
		if pauseCmdZoneId == "" {
			errText = append(errText, "--zone-id")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		response := setting.Pause(pauseCmdZoneId, jsonBody.UnpauseJson)
		if !response.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to unpause Cloudflare Zone. The error is", response.Errors[0].Message)
			os.Exit(1)
		} else {
			output.PrintOut(response.Result, flagQuery, flagOutput)
		}
	},
}

func init() {
	settingCmd.AddCommand(unPauseCmd)
	unPauseCmd.Flags().StringVarP(&pauseCmdZoneId, "zone-id", "", "", "Cloudflare zone ID need to unpaused")
}
