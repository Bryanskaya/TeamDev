package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type DBConfiguration struct {
	Type string `json:"type"`
	Name string `json:"name"`

	User     string `json:"user"`
	Password string `json:"password"`

	Host string `json:"host"`
	Port string `json:"port"`
}

type Configuration struct {
	DB            DBConfiguration `json:"db"`
	LogFile       string          `json:"log_file"`
	Port          uint16          `json:"port"`
	TokenPassword string          `json:"token_password"`
	TokenLiveTime int64           `json:"token_livetime"`
}

var (
	Config Configuration
)

// TODO: returnable error
func InitConfig(args ...string) {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}

	switch len(args) {
	case 1:
		new_port, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		Config.Port = uint16(new_port)
	}

	Config.LogFile = fmt.Sprintf(Config.LogFile, Config.Port)
}
