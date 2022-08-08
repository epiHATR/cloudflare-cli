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

var dnsDeleteCmdZoneId = ""
var dnsDeleteCmdDnsId = ""
var dnsDeleteConfirmText = ""
var dnsForceDelete bool = false

// dnsCmd represents the dns command
var dnsDelete = &cobra.Command{
	Use:   "delete",
	Short: "delete a DNS record to a Cloudflare zone",
	Long:  text.ZoneDndDelete + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(dnsDeleteCmdZoneId) <= 0 {
			errText = append(errText, "--zone-id")
		}

		if len(dnsDeleteCmdDnsId) <= 0 {
			errText = append(errText, "--id|-i")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		res := dns.GetZoneDnsDetail(dnsDeleteCmdZoneId, dnsDeleteCmdDnsId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get DNS record. The error is", res.Errors[0].Message)
			os.Exit(1)
		} else {
			log.Println("DNS record name", res.Result.Name, "with id", res.Result.Id, "will be removed from Cloudflare zone", res.Result.Zone_name)
			if !dnsForceDelete {
				for {
					fmt.Print("Please confirm your action (yes/no): ")
					fmt.Scan(&dnsDeleteConfirmText)
					if strings.ToLower(dnsDeleteConfirmText) == "yes" || strings.ToLower(dnsDeleteConfirmText) == "no" {
						break
					}
				}
			} else {
				dnsDeleteConfirmText = "yes"
			}

			if dnsDeleteConfirmText == "yes" {
				res := dns.DeleteDnsRecord(dnsDeleteCmdZoneId, dnsDeleteCmdDnsId)
				if !res.Success {
					fmt.Fprintln(os.Stderr, "Error: failed to delete DNS record. The error is", res.Errors[0].Message)
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
	dnsCmd.AddCommand(dnsDelete)
	dnsDelete.Flags().StringVarP(&dnsDeleteCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsDelete.Flags().StringVarP(&dnsDeleteCmdDnsId, "id", "i", "", "DNS id need to delete")
	dnsDelete.Flags().BoolVarP(&dnsForceDelete, "--force", "f", false, "force delete dns records without confirm")
}
