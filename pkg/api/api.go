package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type Response struct {
	Success bool    `json:"success"`
	Errors  []Error `json:"errors"`
}

func VerifyToken(input string) (result string) {
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
	return string(responseData)
}

func VerifyKeyEmail(email string, key string) (result string) {
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
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	jsonData, err := json.Marshal(responseObject)
	if err != nil {
		fmt.Println(err)
		return
	}
	return string(jsonData)
}
