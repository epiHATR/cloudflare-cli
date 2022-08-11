package account

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func GetAllAccounts(pageNumber int, queryName string) response.AccountListResponse {
	log.Println("get all Cloudflare managed accounts/organizations")
	queryUrl := endpoint.ApiEndPoint + endpoint.AccountListEndPoint + "&page=" + strconv.Itoa(pageNumber)
	if len(queryName) > 0 {
		queryUrl += "&name=" + queryName
	}

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.AccountListResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}