package util

import (
	"bytes"
	"cloudflare/pkg/structs"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmespath/go-jmespath"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

func LoadConfig() (config *structs.Config) {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".cloudflare" (without extension).
	configFileName := ".cloudflare"
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)
	viper.AutomaticEnv()

	c := &structs.Config{}
	var path = ""

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("configuration file", viper.ConfigFileUsed(), "found")
		if err := viper.Unmarshal(c); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("no configuration file found.")
		// create an empty
		path := filepath.Join(home, configFileName)
		file, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		log.Println("emtpy configuration file created at", path)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.SetEnvPrefix("cf")
	viper.BindEnv("auth.token")
	viper.BindEnv("auth.email")
	viper.BindEnv("auth.key")
	viper.SafeWriteConfigAs(path)
	viper.WriteConfig()
	log.Println("successfully bind environment variables to configuration file values")
	return c
}

func SetToken(token string) (res bool) {
	result := false
	viper.Set("auth.token", token)
	viper.Set("auth.email", "")
	viper.Set("auth.key", "")
	viper.WriteConfig()
	log.Println("saved token", "****************", "to .cloudflare configuration file at ", viper.ConfigFileUsed())
	return result
}

func SetEmailKey(email string, key string) (res bool) {
	result := false
	viper.Set("auth.email", email)
	viper.Set("auth.key", key)
	viper.Set("auth.token", "")
	viper.WriteConfig()
	log.Println("saved email & key to .cloudflare configuration file at ", viper.ConfigFileUsed())
	return result
}

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
