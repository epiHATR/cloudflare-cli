/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/plan"
	"cloudflare/pkg/consts/jsonBody"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var updatePlanCmdZoneId = ""
var updatePlanCmdPlanId = ""
var updatePlanCmdPlanName = ""

// upgradeCmd represents the update command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade Cloudflare zone's plan",
	Long:  text.UpdatePlanLongtext + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}

		if len(updatePlanCmdZoneId) == 0 {
			errText = append(errText, "--zone-id")
		}

		if len(updatePlanCmdPlanId) <= 0 {
			if len(updatePlanCmdPlanName) <= 0 {
				log.Println("plan id")
				errText = append(errText, "--plan-id or --plan-name")
			}
			log.Println("plan name", updatePlanCmdPlanName)
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		id := ""

		if len(updatePlanCmdPlanId) > 0 {
			id = updatePlanCmdPlanId
		} else {
			res := plan.ListAllAvailablePlan(updatePlanCmdZoneId)
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to get all available plans. The error is", res.Errors[0].Message)
				os.Exit(1)
			} else {
				log.Println("found", len(res.Result), "plan")
				for _, planItem := range res.Result {
					log.Println(planItem.Name)
					if strings.ToLower(planItem.Legacy_id) == strings.ToLower(updatePlanCmdPlanName) {
						id = planItem.Id
						break
					}
				}
			}
			if len(id) <= 0 {
				fmt.Fprintln(os.Stderr, "Error: no available plan called", updatePlanCmdPlanName, text.SubCmdHelpText)
				os.Exit(1)
			}
		}

		res := plan.SetPlan(updatePlanCmdZoneId, fmt.Sprintf(jsonBody.PlanType, id))
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to upgrade plan for zone. The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	zonePlanCmd.AddCommand(upgradeCmd)
	upgradeCmd.Flags().StringVarP(&updatePlanCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	upgradeCmd.Flags().StringVarP(&updatePlanCmdPlanId, "plan-id", "i", "", "plan ID")
	upgradeCmd.Flags().StringVarP(&updatePlanCmdPlanName, "plan-name", "n", "", "plan ID")
}
