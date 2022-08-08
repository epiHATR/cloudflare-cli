package plan

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ListAllAvailablePlan(zoneId string) response.ZonePlanResponse {
	log.Println("getting all available plan for zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AvailablePlanEndpoint, zoneId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.ZonePlanResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func AvailablePlanDetail(zoneId string, planId string) response.ZonePlanResponse {
	log.Println("getting all available plan for zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AvailablePlanEndpoint+"/"+planId, zoneId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.ZonePlanResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func ListAllRatePlan(zoneId string) response.ZonePlanResponse {
	log.Println("getting all available plan for zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AvailableRatePlanEndpoint, zoneId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.ZonePlanResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func SetPlan(zoneId string, planJson string) response.ZoneDetailResponse {
	log.Println("UNPAUSE: changing zone setting on zoneId", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "PATCH", planJson)
	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
