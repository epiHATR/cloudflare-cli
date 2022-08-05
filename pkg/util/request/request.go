package request

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func CreateRequest(url string, method string, body io.Reader) []byte {
	log.Println("query url: ", url)
	log.Println("method: ", method)

	req, err := http.NewRequest(method, url, body)
	token := viper.GetString("auth.token")
	if token != "" {
		log.Println("performing HTTP request using API Token")
		req.Header.Add("Authorization", "Bearer "+token)
	} else {
		email := viper.GetString("auth.email")
		key := viper.GetString("auth.key")
		if email != "" && key != "" {
			log.Println("performing HTTP request using X-Auth-Email & X-Auth-Key")
			req.Header.Add("X-Auth-Email", email)
			req.Header.Add("X-Auth-Key", key)
		} else {
			fmt.Fprintln(os.Stderr, "Error: No valid credential provided, please run `cloudflare login --help` for instructions")
			os.Exit(1)
		}
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: Failed to perform HTTP request to server. The error was ", err.Error())
		os.Exit(1)
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read response data. The error was ", err.Error())
		os.Exit(1)
	}
	return resData
}
