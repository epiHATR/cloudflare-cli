package config

import (
	"cloudflare/pkg/model/config"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LoadConfig() (conf *config.Config) {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".cloudflare" (without extension).
	configFileName := ".cloudflare"
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(configFileName)
	viper.AutomaticEnv()

	c := &config.Config{}
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
	if len(token) <= 0 {
		return result
	}

	viper.Set("auth.token", token)
	viper.Set("auth.email", "")
	viper.Set("auth.key", "")
	viper.WriteConfig()
	result = true
	log.Println("saved token", "****************", "to .cloudflare configuration file at ", viper.ConfigFileUsed())
	return result
}

func SetEmailKey(email string, key string) (res bool) {
	result := false

	if len(email) > 0 && len(key) > 0 && strings.Contains(email, "@") == true {
		viper.Set("auth.email", email)
		viper.Set("auth.key", key)
		viper.Set("auth.token", "")
		viper.WriteConfig()
		log.Println("saved email & key to .cloudflare configuration file at ", viper.ConfigFileUsed())
		result = true
	} else {
		result = false
	}
	return result
}
