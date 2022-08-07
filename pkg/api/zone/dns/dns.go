package dns

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetZoneDns(zoneId string, recordType string) response.DnsListResponse {
	log.Println("getting DNS records in zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.ZoneDnsListEndpoint, zoneId)
	if recordType != "" {
		queryUrl += "&type=" + recordType
	}

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.DnsListResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func GetZoneDnsDetail(zoneId string, dnsId string) response.DnsDetailResponse {
	log.Println("getting DNS details in zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.ZoneDetails+"/"+dnsId, zoneId)
	respData := request.CreateRequest(queryUrl, "GET", "")

	resObj := response.DnsDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func AddDNSRecord(zoneId string, jsonData string) response.DnsDetailResponse {
	log.Println("creating DNS record in", zoneId)

	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.ZoneDetails, zoneId)
	respData := request.CreateRequest(queryUrl, "POST", jsonData)

	resObj := response.DnsDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
