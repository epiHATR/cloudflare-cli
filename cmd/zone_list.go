/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"bytes"
	"cloudflare/pkg/text"
	"cloudflare/pkg/util"
	methods "cloudflare/pkg/zone"
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

var accountId = ""
var name = ""

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all Cloudflare zones",
	Long:  text.ZoneCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {

			response := methods.GetAllZone(1, accountId, name)

			if !response.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to list Cloudflare zones. The error is ", response.Errors[0].Message)
				os.Exit(1)
			} else {
				log.Println("number of page: ", response.Result_Info.Total_pages)
				switch output {

				case "json":
					fmt.Println(util.ToPrettyJson(response.Result, query))

				case "yaml":
					fmt.Println(util.ToPrettyYaml(response.Result, query))

				default:
					fmt.Println(util.ToPrettyJson(response.Result, query))
				}
			}
		}
	},
}

func init() {
	zoneCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&accountId, "account-id", "", "", "cloudflare account id (organization id)")
	listCmd.Flags().StringVarP(&name, "name", "n", "", "specify zone name to search")
}
