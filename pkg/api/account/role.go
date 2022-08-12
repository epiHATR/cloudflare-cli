package account

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetAccountRoles(accountId string) response.AccountRolesResponse {
	log.Println("get all roles in a Cloudflare managed account/organization")
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AccountRolesEndpoint, accountId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.AccountRolesResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func GetAccountRoleDetails(accountId string, roleId string) response.AccountRoleDetailResponse {
	log.Println("get all roles in a Cloudflare managed account/organization")
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AccountRolesEndpoint+"/"+roleId, accountId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.AccountRoleDetailResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
