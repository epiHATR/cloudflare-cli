package cache

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"cloudflare/pkg/util/request"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func PurgeEverything(zoneId string, payload string) response.ZoneDetailResponse {
	log.Println("caches purge everything in zone", zoneId)
	queryUrl := fmt.Sprintf(endpoint.ApiEndPoint+endpoint.ZoneCachePurge, zoneId)

	respData := request.CreateRequest(queryUrl, "POST", payload)

	resObj := response.ZoneDetailResponse{}
	err := json.Unmarshal(respData, &resObj)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return resObj
}
