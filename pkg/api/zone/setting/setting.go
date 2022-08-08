package setting

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"log"
	"os"
)

func Pause(zoneId string, settingJson string) response.ZoneDetailResponse {
	log.Println("PAUSE: changing zone setting on zoneId", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "PATCH", settingJson)
	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func UnPause(zoneId string, settingJson string) response.ZoneDetailResponse {
	log.Println("UNPAUSE: changing zone setting on zoneId", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "PATCH", settingJson)
	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}

func SetType(zoneId string, zoneTypeJson string) response.ZoneDetailResponse {
	log.Println("UNPAUSE: changing zone setting on zoneId", zoneId)
	queryUrl := endpoint.ApiEndPoint + endpoint.ZoneDetailEndPoint + "/" + zoneId
	respData := request.CreateRequest(queryUrl, "PATCH", zoneTypeJson)
	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
