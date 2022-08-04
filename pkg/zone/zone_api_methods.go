package zone_api_methods

import (
	"cloudflare/pkg/consts"
	"cloudflare/pkg/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func GetAllZone() (result structs.ZoneListResponse) {
	log.Println("get all zone in all Cloudflare Accounts")

	client := &http.Client{}
	req, err := http.NewRequest("GET", consts.ApiEndPoint+consts.ZoneListEndPoint, nil)
	req.Header.Add("Authorization", "Bearer "+viper.GetString("auth.token"))
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := structs.ZoneListResponse{}
	_ = json.Unmarshal(responseData, &res)
	return res

}

func GetZoneById(id string) {
	log.Println("get zone by ID")
}
