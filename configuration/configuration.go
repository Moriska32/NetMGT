package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Configuration - Общая конфигурация сервера.
type Configuration struct {
	PostgresDatabaseCfg PostgresDatabaseCfg `json:"postgres_database_cfg"`
	ServerCfg           struct {
		Port    string `json:"port"`
		PortSSL string `json:"portSSL"`
	} `json:"server_cfg"`
	RedisCfg RedisDatabaseCfg `json:"redis_cfg"`
}

// SetParams - инициализация параметров сервера
func (cfg *Configuration) SetParams(fname string) {

	configFile, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Println("Error reading configuration file:", err)
		return
	}
	err = json.Unmarshal(configFile, &cfg)
	if err != nil {
		fmt.Println("Error parsing configuration data:", err)
		return
	}

	fmt.Println("Using next configuration:")
	fmt.Println("\t Postgres Database:")
	for i := range (*cfg).PostgresDatabaseCfg {
		fmt.Println("\t\tDB index:", i)
		fmt.Println("\t\tHost:", (*cfg).PostgresDatabaseCfg[i].Host)
		fmt.Println("\t\tPort:", (*cfg).PostgresDatabaseCfg[i].Port)
		fmt.Println("\t\tDatabase name:", (*cfg).PostgresDatabaseCfg[i].Database)
		fmt.Println("\t\tContext database name:", (*cfg).PostgresDatabaseCfg[i].Name)
		fmt.Println("\t\tUser", (*cfg).PostgresDatabaseCfg[i].User)
		fmt.Println("\t\tPassword", (*cfg).PostgresDatabaseCfg[i].Password)
		fmt.Println("\t\tEnable TLS:", (*cfg).PostgresDatabaseCfg[i].EnableTLS)
	}
	fmt.Println("\tServer:")
	fmt.Println("\t\tServer port:", (*cfg).ServerCfg.Port)
}
