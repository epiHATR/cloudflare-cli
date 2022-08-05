package authenticate

import (
	"cloudflare/pkg/consts/endpoint"
	"cloudflare/pkg/model/response"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func VerifyToken(input string) (result response.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint.ApiEndPoint+endpoint.TokenVerifyEndPoint, nil)
	req.Header.Add("Authorization", "Bearer "+input)
	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	respData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp := response.Response{}
	_ = json.Unmarshal(respData, &resp)
	return resp
}

func VerifyKeyEmail(email string, key string) (result response.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpoint.ApiEndPoint+endpoint.EmailKeyVerifyEndPoint, nil)
	req.Header.Add("X-Auth-Email", email)
	req.Header.Add("X-Auth-Key", key)
	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp := response.Response{}
	_ = json.Unmarshal(respData, &resp)
	return resp
}
