package database

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rotisserie/eris"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Определение структуры для конфигурации
type DBConfig struct {
	Logins struct {
		Main struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		} `json:"main"`
	} `json:"logins"`
	ServerType   string `json:"server_type"`
	ServerAddres string `json:"server_addres"`
}

var Db *gorm.DB

func InitDB() error {
	config, er := getDBConfig("config.json")
	if er != nil {
		eris.Wrap(er, "failed to open config file")
	}

	connect := fmt.Sprintf("%v%v:%v%v",
		config.ServerType,
		config.Logins.Main.Login,
		config.Logins.Main.Password,
		config.ServerAddres)

	db, er := gorm.Open(sqlserver.Open(connect), &gorm.Config{})
	if er != nil {
		return eris.Wrap(er, "failed to connect to database")
	}
	Db = db
	fmt.Print("data base is connected\n")
	return nil
}

func getDBConfig(configName string) (*DBConfig, error) {
	file, er := os.Open(configName)
	if er != nil {
		return nil, er
	}
	defer file.Close()

	var dbConfig *DBConfig
	er = json.NewDecoder(file).Decode(&dbConfig)
	if er != nil {
		return nil, er
	}
	return dbConfig, nil
}
