/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var zoneShowFlagZoneName = ""
var zoneShowFlagZoneId = ""

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show a cloudflare zone details",
	Long:  text.ZoneShowCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		//run function based on input flag'
		if zoneShowFlagZoneId != "" {
			res := response.ZoneDetailResponse{}
			res = zone.GetZoneById(zoneShowFlagZoneId)
			//manipulate response
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to get Cloudflare zone details. The error is", res.Errors[0].Message)
				os.Exit(1)
			} else {
				output.PrintOut(res.Result, flagQuery, flagOutput)
			}

		} else if zoneShowFlagZoneName != "" {
			res := response.ZoneListResponse{}
			res = zone.GetZoneByName(zoneShowFlagZoneName)
			//manipulate response
			if !res.Success {
				fmt.Fprintln(os.Stderr, "Error: failed to get Cloudflare zone details. The error is", res.Errors[0].Message)
				os.Exit(1)
			} else {
				output.PrintOut(res.Result[0], flagQuery, flagOutput)
			}

		} else {
			fmt.Fprintln(os.Stderr, "Error: please specify at least one of following value --name|-n, --id|-i")
			os.Exit(1)
		}

	},
}

func init() {
	zoneCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&zoneShowFlagZoneName, "name", "n", "", "name of Cloudflare zone")
	showCmd.Flags().StringVarP(&zoneShowFlagZoneId, "id", "i", "", "id of Cloudflare zone")
}
