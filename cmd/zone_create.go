/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone"
	"cloudflare/pkg/api/zone/plan"
	"cloudflare/pkg/consts/jsonBody"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var createCmdZoneName string = ""
var createCmdZoneType string = ""
var createCmdAccountId string = ""
var createCmdJumpStart bool = true
var createCmdPlanName string = "free"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new Cloudflare zone",
	Long:  text.ZoneCreateLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(createCmdZoneName) <= 0 {
			errText = append(errText, "--name")
		}
		if len(createCmdAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		if strings.ToLower(createCmdZoneType) != "partial" && strings.ToLower(createCmdZoneType) != "full" {
			fmt.Fprintln(os.Stderr, "Error: Invalid data for flag --type, value must be in 'partial, full'.")
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		res := zone.CreateNewZone(createCmdZoneName, createCmdAccountId, createCmdZoneType, createCmdJumpStart)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: Failed to create Zone", createCmdZoneName, ". The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		planId := plan.GetZonePlanId(res.Result.Id, createCmdPlanName)
		if len(planId) <= 0 {
			fmt.Fprintln(os.Stderr, "Error: Failed to get plan ", createCmdZoneName, "details. The error is", res.Errors[0].Message)
			os.Exit(1)
		}

		res = plan.SetPlan(res.Result.Id, fmt.Sprintf(jsonBody.PlanType, planId))
		if !res.Success {
			//we're going to remove failed zone here
			fmt.Fprintln(os.Stderr, "Error: Failed to upgrade", createCmdZoneName, "to plan", createCmdPlanName, ". The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	zoneCmd.AddCommand(createCmd)
	//createCmd.Flags().SortFlags = false
	createCmd.Flags().StringVarP(&createCmdZoneName, "name", "n", "", "Name of the zone need to be created.")
	createCmd.Flags().StringVarP(&createCmdAccountId, "account-id", "", "", "Id of Cloudflare Organization/Account which zone is being created in.")
	createCmd.Flags().StringVarP(&createCmdZoneType, "type", "t", "partial", "Type of zone, must be one of 'partial, full'")
	createCmd.Flags().StringVarP(&createCmdPlanName, "plan-name", "", "free", "Type of plan which the zone is being created on, must be one of 'free, pro, enterprise, business'")
	createCmd.Flags().BoolVarP(&createCmdJumpStart, "fetch-existing-dns", "", true, "Automatically attempt to fetch existing DNS records")
}
