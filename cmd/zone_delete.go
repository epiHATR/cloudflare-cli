/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone"
	"cloudflare/pkg/api/zone/plan"
	"cloudflare/pkg/consts/jsonBody"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var deleteZoneCmdZoneId string = ""
var deleteZoneCmdForceDelete bool = false
var deleteZoneCmdConfirmationText string = ""

var deleteZoneCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a Cloudflare zone",
	Long:  text.DeleteZoneLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if len(deleteZoneCmdZoneId) <= 0 {
			errText = append(errText, "--zone-id")
		}
		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := zone.GetZoneById(deleteZoneCmdZoneId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get Cloudflare zone. The error is", res.Errors[0].Message)
			os.Exit(1)
		} else {
			log.Println("Zone name", res.Result.Name, "with id(", res.Result.Id, ") is being deleted from Cloudflare")
			if !deleteZoneCmdForceDelete {
				for {
					fmt.Print("Delete: ", res.Result.Name, "with id", res.Result.Id, ". Please confirm your action (yes/no): ")
					fmt.Scan(&deleteZoneCmdConfirmationText)
					if strings.ToLower(deleteZoneCmdConfirmationText) == "yes" || strings.ToLower(deleteZoneCmdConfirmationText) == "no" {
						break
					}
				}
			} else {
				deleteZoneCmdConfirmationText = "yes"
			}

			if deleteZoneCmdConfirmationText == "yes" {
				planId := plan.GetZonePlanId(res.Result.Id, "free")
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

				res = zone.DeleteAZone(res.Result.Id)
				if !res.Success {
					//we're going to remove failed zone here
					fmt.Fprintln(os.Stderr, "Error: Failed to delete zone", deleteZoneCmdZoneId, ". The error is", res.Errors[0].Message)
					os.Exit(1)
				}
				result := response.CmdResponse{}
				result.Success = res.Success
				output.PrintOut(result, flagQuery, flagOutput)
			}
		}
	},
}

func init() {
	zoneCmd.AddCommand(deleteZoneCmd)
	deleteZoneCmd.Flags().StringVarP(&deleteZoneCmdZoneId, "zone-id", "", "", "Id of a Cloudflare zone need to be deleted.")
	deleteZoneCmd.Flags().BoolVarP(&deleteZoneCmdForceDelete, "force", "f", false, "Force delete without confirmation")
}
