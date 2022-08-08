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

var planShowCmdZoneId string = ""
var planShowCmdPlanId string = ""

// showCmd represents the show command
var planShow = &cobra.Command{
	Use:   "show",
	Short: "show Cloudflare plan details",
	Long:  text.PlanDetailsLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(planShowCmdZoneId) <= 0 {
			errText = append(errText, "--zone-id")
		}
		if len(planShowCmdPlanId) <= 0 {
			errText = append(errText, "--id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		res := plan.AvailablePlanDetail(planShowCmdZoneId, planShowCmdPlanId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get plan details for zone. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	zonePlanCmd.AddCommand(planShow)
	planShow.Flags().StringVarP(&planShowCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	planShow.Flags().StringVarP(&planShowCmdPlanId, "id", "i", "", "plan id")
}
