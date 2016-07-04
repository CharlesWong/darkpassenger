package account

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var (
	config *Config
)

type Config struct {
	AdminToken string
	ListenAddr string
	DataFile   string
}

func newConfig(file string) (*Config, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, errors.New("Error reading config file.")
	}

	c := &Config{}
	err = json.Unmarshal(b, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func InitConfig(file string) (err error) {
	config, err = newConfig(file)
	return
}

func GetConfig() *Config {
	return config
}
