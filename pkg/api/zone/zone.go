package zone

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetAllZone(pageNumber int, account string, searchKey string) (result response.ZoneListResponse) {
	log.Println("get all zone in Cloudflare account")

	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneListEndPoint + "&page=" + strconv.Itoa(pageNumber)
	if account != "" {
		queryUrl += "&account.id=" + account
	}
	if searchKey != "" {
		queryUrl += "&name=" + searchKey
	}

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.ZoneListResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj

}

func GetZoneById(zoneId string) response.ZoneDetailResponse {
	log.Println("get zone by ID", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "GET", "")

	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func GetZoneByName(name string) response.ZoneListResponse {
	log.Println("get zone by name", name)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "?name=" + name
	respData := request.CreateRequest(queryUrl, "GET", "")

	resObj := response.ZoneListResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func CreateNewZone(name string, accountId string, zoneType string, jump_start bool) response.ZoneDetailResponse {
	if len(name) <= 0 || len(accountId) <= 0 {
		os.Exit(1)
	}
	log.Println("creating zone ", name, "in account", accountId)
	queryUrl := endpoint.ApiEndPoint + endpoint.CreateZoneEndpoint
	jsonBody := fmt.Sprintf(`{ "name":"%s","account": { "id":"%s"}, "type": "%s", "jump_start": %v }`, name, accountId, zoneType, jump_start)
	log.Println(jsonBody)
	respData := request.CreateRequest(queryUrl, "POST", jsonBody)

	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func DeleteAZone(zoneId string) response.ZoneDetailResponse {
	if len(zoneId) <= 0 || len(zoneId) <= 0 {
		os.Exit(1)
	}
	log.Println("deleting zone ", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.DeleteZoneEndpoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "DELETE", "")

	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
