package api

import (
	"cloudflare/pkg/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func VerifyToken(input string) (result structs.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/user/tokens/verify", nil)
	req.Header.Add("Authorization", "Bearer "+input)
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := structs.Response{}
	_ = json.Unmarshal(responseData, &res)
	return res
}

func VerifyKeyEmail(email string, key string) (result structs.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.cloudflare.com/client/v4/user", nil)
	req.Header.Add("X-Auth-Email", email)
	req.Header.Add("X-Auth-Key", key)
	response, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	res := structs.Response{}
	_ = json.Unmarshal(responseData, &res)
	return res
}
