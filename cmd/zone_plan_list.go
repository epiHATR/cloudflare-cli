/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/plan"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var availablePlans bool = false
var ratePlanOnly bool = false

var availablePlanZoneId string = ""

// listCmd represents the list command
var zonePlanListCmd = &cobra.Command{
	Use:   "list",
	Short: "list plans for the Cloudflare zone",
	Long:  text.AvailablePlanLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if availablePlanZoneId == "" {
			errText = append(errText, "--zone-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		if availablePlans {
			res := plan.ListAllAvailablePlan(availablePlanZoneId)
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to get all available plan for zone. The error is", res.Errors[0].Message)
				os.Exit(1)
			}
			output.PrintOut(res.Result, flagQuery, flagOutput)
		} else if ratePlanOnly {
			res := plan.ListAllRatePlan(availablePlanZoneId)
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to get all rate plan for zone. The error is", res.Errors[0].Message)
				os.Exit(1)
			}
			output.PrintOut(res.Result, flagQuery, flagOutput)
		} else {
			fmt.Fprintln(os.Stderr, "Error: You need to specify at least on of following arguments --all-available|-a, --rate-plan-only ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
	},
}

func init() {
	zonePlanCmd.AddCommand(zonePlanListCmd)
	zonePlanListCmd.Flags().StringVarP(&availablePlanZoneId, "zone-id", "", "", "cloudlfare zone ID")
	zonePlanListCmd.Flags().BoolVarP(&availablePlans, "all-available", "a", false, "list available plans the zone can subscribe to")
	zonePlanListCmd.Flags().BoolVarP(&ratePlanOnly, "rate-plan-only", "", false, "list availabe rate plan")
}
