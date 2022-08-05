package output

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmespath/go-jmespath"
	"gopkg.in/yaml.v2"
)

func ToPrettyJson(input interface{}, query string) string {
	var result interface{}
	var err error

	if query != "" {
		result, err = jmespath.Search(query, input)
	} else {
		result = input
	}

	b, err := json.Marshal(result)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to marshal object ", err.Error())
		os.Exit(1)
	}

	var finalJSONString bytes.Buffer
	if err := json.Indent(&finalJSONString, []byte(b), "", "    "); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return finalJSONString.String()
}

func ToPureJson(input interface{}, query string) string {

	var result interface{}
	var err error

	if query != "" {
		result, err = jmespath.Search(query, input)
	} else {
		result = input
	}
	b, err := json.Marshal(result)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to marshal object ", err.Error())
		os.Exit(1)
	}
	return string(b)
}

func ToPrettyYaml(input interface{}, query string) string {
	var result interface{}
	var err error

	if query != "" {
		result, err = jmespath.Search(query, input)
	} else {
		result = input
	}
	y, err := yaml.Marshal(result)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	return string(y)
}
