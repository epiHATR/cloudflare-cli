/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/dns"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var dnsUpdateCmdZoneId = ""
var dnsUpdateCmdDnsId = ""
var dnsUpdateCmdData = ""
var dnsForceUpdate bool = false
var dnsUpdateConfirmText = ""

// dnsCmd represents the dns command
var dnsUpdate = &cobra.Command{
	Use:   "update",
	Short: "update a DNS record to a Cloudflare zone",
	Long:  text.ZoneDnsUpdate + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(dnsUpdateCmdZoneId) <= 0 {
			errText = append(errText, "--zone-id")
		}
		if len(dnsUpdateCmdDnsId) <= 0 {
			errText = append(errText, "--id|-i")
		}
		if len(dnsUpdateCmdData) <= 0 {
			errText = append(errText, "--data|-d")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := dns.GetZoneDnsDetail(dnsUpdateCmdZoneId, dnsUpdateCmdDnsId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get DNS record. The error is", res.Errors[0].Message)
			os.Exit(1)
		} else {
			log.Println("DNS record name", res.Result.Name, "with id", res.Result.Id, "will be updated on Cloudflare zone", res.Result.Zone_name)
			if !dnsForceUpdate {
				for {
					fmt.Print("Please confirm your action (yes/no): ")
					fmt.Scan(&dnsUpdateConfirmText)
					if strings.ToLower(dnsUpdateConfirmText) == "yes" || strings.ToLower(dnsUpdateConfirmText) == "no" {
						break
					}
				}
			} else {
				dnsUpdateConfirmText = "yes"
			}
			if dnsUpdateConfirmText == "yes" {
				res := dns.UpdateDnsRecord(dnsUpdateCmdZoneId, dnsUpdateCmdDnsId, dnsUpdateCmdData)
				if !res.Success {
					fmt.Fprintln(os.Stderr, "Error: failed to update DNS record. The error is", res.Errors[0].Message)
					os.Exit(1)
				} else {
					result := response.CmdResponse{}
					result.Success = res.Success
					output.PrintOut(result, flagQuery, flagOutput)
				}
			}
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsUpdate)
	dnsUpdate.Flags().StringVarP(&dnsUpdateCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsUpdate.Flags().StringVarP(&dnsUpdateCmdDnsId, "id", "i", "", "DNS id need to update")
	dnsUpdate.Flags().StringVarP(&dnsUpdateCmdData, "data", "d", "", "DNS id need to update")
	dnsUpdate.Flags().BoolVarP(&dnsForceUpdate, "--force", "f", false, "force update DNS records without confirm")
}
