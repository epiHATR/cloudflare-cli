/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"bytes"
	"cloudflare/pkg/api/zone"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"encoding/json"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

var listCmdAccountId = ""
var listCmdAccountName = ""

// zoneListCmd represents the list command
var zoneListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all Cloudflare zones",
	Long:  text.ZoneListCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			result := []response.Result{}

			response := zone.GetAllZone(1, listCmdAccountId, listCmdAccountName)
			bar := progressbar.Default(int64(response.Result_Info.Total_pages))
			if !response.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is", response.Errors[0].Message)
				fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
				os.Exit(1)
			} else {
				result = append(result, response.Result...)
				bar.Add(1)
				if response.Result_Info.Total_pages >= 2 {
					for i := 2; i <= response.Result_Info.Total_pages; i++ {
						bar.Add(1)
						//fmt.Fprintln(os.Stdin, "Fetching cloudflare zone page", i, "/", response.Result_Info.Total_pages)
						res := zone.GetAllZone(i, listCmdAccountId, listCmdAccountName)
						if !res.Success {
							fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is", response.Errors[0].Message)
							fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
							os.Exit(1)
						} else {
							result = append(result, res.Result...)
						}
					}
				}

				if len(flagQuery) <= 0 {
					flagQuery = "[].{id:id, name:name, status:status, account:{id: account.id, name: account.name}}"
				}
				output.PrintOut(result, flagQuery, flagOutput)
			}
		}
	},
}

func init() {
	zoneCmd.AddCommand(zoneListCmd)
	zoneListCmd.Flags().StringVarP(&listCmdAccountId, "account-id", "", "", "cloudflare account id (organization id)")
	zoneListCmd.Flags().StringVarP(&listCmdAccountName, "account-name", "n", "", "specify zone name to search")

}
