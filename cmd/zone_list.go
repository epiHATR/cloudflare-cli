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

var listCmdAccountId = ""
var listCmdAccountName = ""

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all Cloudflare zones",
	Long:  text.ZoneListCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			result := []response.Result{}

			response := zone.GetAllZone(1, listCmdAccountId, listCmdAccountName)
			if !response.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is", response.Errors[0].Message)
				fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
				os.Exit(1)
			} else {
				log.Println("number of page: ", response.Result_Info.Total_pages)
				result = append(result, response.Result...)
				if response.Result_Info.Total_pages > 2 && response.Result_Info.Total_pages <= 10 {
					for i := 2; i < response.Result_Info.Total_pages; i++ {
						response := zone.GetAllZone(i, listCmdAccountId, listCmdAccountName)
						if !response.Success {
							fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is", response.Errors[0].Message)
							fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
							os.Exit(1)
						} else {
							result = append(result, response.Result...)
						}
					}
				}

				if len(flagQuery) <= 0 {
					flagQuery = "[].{id:id, name:name, status:status, account:{id: account.id, name: account.name}}"
				}
				output.PrintOut(result, flagQuery, flagOutput)
				if response.Result_Info.Total_pages > 10 {
					fmt.Fprintf(os.Stdin, fmt.Sprintf("Too many result returned, displaying %d/%d (page %d/%d)", response.Result_Info.Count, response.Result_Info.Total_count, response.Result_Info.Page, response.Result_Info.Total_pages))
					fmt.Fprintf(os.Stdin, fmt.Sprintln())
				}
			}
		}
	},
}

func init() {
	zoneCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listCmdAccountId, "account-id", "", "", "cloudflare account id (organization id)")
	listCmd.Flags().StringVarP(&listCmdAccountName, "account-name", "n", "", "specify zone name to search")

}
