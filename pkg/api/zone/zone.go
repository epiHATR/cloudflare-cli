package zone

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
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

	respData := request.CreateRequest(queryUrl, "GET", nil)
	resObj := response.ZoneListResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj

}

func GetZoneById(id string) response.ZoneDetailResponse {
	log.Println("get zone by ID", id)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + id
	respData := request.CreateRequest(queryUrl, "GET", nil)

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

	log.Println("the request url: ", queryUrl)
	respData := request.CreateRequest(queryUrl, "GET", nil)

	resObj := response.ZoneListResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
