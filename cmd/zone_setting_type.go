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

var setTypeCmdZoneId string = ""
var setTypeCmdType string = ""

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set-type",
	Short: "change Cloudflare zone type",
	Long:  text.SetTypeLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("verify flag inputs")
		errText := []string{}
		if setTypeCmdZoneId == "" {
			errText = append(errText, "--zone-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		zoneTypeJson := ""
		if setTypeCmdType == "full" {
			zoneTypeJson = jsonBody.FullType
		} else if setTypeCmdType == "partial" {
			zoneTypeJson = jsonBody.PartialType
		}

		response := setting.SetType(setTypeCmdZoneId, zoneTypeJson)
		if !response.Success {
			fmt.Fprintln(os.Stderr, "Error: failed change Cloudflare zone type. The error is", response.Errors[0].Message)
			os.Exit(1)
		} else {
			output.PrintOut(response.Result, flagQuery, flagOutput)
		}
	},
}

func init() {
	settingCmd.AddCommand(setCmd)
	setCmd.Flags().StringVarP(&setTypeCmdZoneId, "zone-id", "", "", "Cloudflare zone ID need to update")
	setCmd.Flags().StringVarP(&setTypeCmdType, "type", "t", "", "Zone type: partial, full")
}
