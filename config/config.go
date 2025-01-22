package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	MAL *struct {
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
	} `json:"mal"`
}

func Load(file string) (*Config, error) {
	conf := &Config{}
	f1, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f1.Close()
	err = json.NewDecoder(f1).Decode(conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
