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

var zoneCreateZoneName string = ""
var zoneCreateZoneType string = ""
var zoneCreateAccountId string = ""
var zoneCreateJumpStart bool = true
var zoneCreatePlanName string = "free"

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create new Cloudflare zone",
	Long:  text.ZoneCreateLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(zoneCreateZoneName) <= 0 {
			errText = append(errText, "--name")
		}
		if len(zoneCreateAccountId) <= 0 {
			errText = append(errText, "--account-id")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		if strings.ToLower(zoneCreateZoneType) != "partial" && strings.ToLower(zoneCreateZoneType) != "full" {
			fmt.Fprintln(os.Stderr, "Error: Invalid data for flag --type, value must be in 'partial, full'.")
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		res := zone.CreateNewZone(zoneCreateZoneName, zoneCreateAccountId, zoneCreateZoneType, zoneCreateJumpStart)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: Failed to create Zone", zoneCreateZoneName, ". The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		planId := plan.GetZonePlanId(res.Result.Id, zoneCreatePlanName)
		if len(planId) <= 0 {
			fmt.Fprintln(os.Stderr, "Error: Failed to get plan ", zoneCreateZoneName, "details. The error is", res.Errors[0].Message)
			os.Exit(1)
		}

		res = plan.SetPlan(res.Result.Id, fmt.Sprintf(jsonBody.PlanType, planId))
		if !res.Success {
			//we're going to remove failed zone here
			fmt.Fprintln(os.Stderr, "Error: Failed to upgrade", zoneCreateZoneName, "to plan", zoneCreatePlanName, ". The error is", res.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(res.Result, flagQuery, flagOutput)
	},
}

func init() {
	zoneCmd.AddCommand(createCmd)
	//createCmd.Flags().SortFlags = false
	createCmd.Flags().StringVarP(&zoneCreateZoneName, "name", "n", "", "Name of the zone need to be created.")
	createCmd.Flags().StringVarP(&zoneCreateAccountId, "account-id", "", "", "Id of Cloudflare Organization/Account which zone is being created in.")
	createCmd.Flags().StringVarP(&zoneCreateZoneType, "type", "t", "partial", "Type of zone, must be one of 'partial, full'")
	createCmd.Flags().StringVarP(&zoneCreatePlanName, "plan-name", "", "free", "Type of plan which the zone is being created on, must be one of 'free, pro, enterprise, business'")
	createCmd.Flags().BoolVarP(&zoneCreateJumpStart, "fetch-existing-dns", "", true, "Automatically attempt to fetch existing DNS records")
}
