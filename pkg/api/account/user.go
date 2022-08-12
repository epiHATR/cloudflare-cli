package account

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/payload"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func GetAccountUsers(accountId string) response.AccountUsersResponse {
	log.Println("get all user in Cloudflare managed accounts/organizations")
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AccountUsersEndpoint, accountId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.AccountUsersResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func GetAccountUserDetail(accountId string, userId string) response.AccountUserDetailResponse {
	log.Println("get details of an users in a Cloudflare managed accounts/organizations")
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AccountUsersEndpoint+"/"+userId, accountId)

	respData := request.CreateRequest(queryUrl, "GET", "")
	resObj := response.AccountUserDetailResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func AccountAddUser(accountId string, email string, status string, roles []string) response.AccountUserDetailResponse {
	log.Println("get details of an users in a Cloudflare managed accounts/organizations")
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.AccountUsersEndpoint, accountId)

	reqBody := payload.UserAddRequest{}
	reqBody.Email = email
	reqBody.Status = status
	reqBody.Roles = roles

	jsonPayload, _ := json.Marshal(reqBody)
	log.Println(string(jsonPayload))

	respData := request.CreateRequest(queryUrl, "POST", string(jsonPayload))
	resObj := response.AccountUserDetailResponse{}

	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
