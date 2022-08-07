/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"bytes"
	"cloudflare/pkg/api/zone"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

var zoneListFlagAccountId = ""
var zoneListFlagAccountName = ""

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all Cloudflare zones",
	Long:  text.ZoneListCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			response := zone.GetAllZone(1, zoneListFlagAccountId, zoneListFlagAccountName)
			if !response.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is", response.Errors[0].Message)
				fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
				os.Exit(1)
			} else {
				log.Println("number of page: ", response.Result_Info.Total_pages)
				output.PrintOut(response.Result, flagQuery, flagOutput)
			}
		}
	},
}

func init() {
	zoneCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&zoneListFlagAccountId, "account-id", "", "", "cloudflare account id (organization id)")
	listCmd.Flags().StringVarP(&zoneListFlagAccountName, "account-name", "n", "", "specify zone name to search")
}
