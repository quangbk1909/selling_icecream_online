package database

import (
	"fmt"
	"log"
	"os"
	"vinid_project/utility"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Address  string `json:"address"`
}
type Config struct {
	Database `json:"database"`
}

type Dao struct {
	db *gorm.DB
}

func NewDao(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}

func SetConfig() Config {
	//doc file config local
	configFile, err := os.Open("./config.json")
	if err != nil {
		// neu khong thay file config local doc file config default
		configFile, err = os.Open("./config.default.json")
		if err != nil {
			panic(err)
		}
		defer configFile.Close()
	}
	defer configFile.Close()

	var conf Config
	err = utility.DecodeDataFromJsonFile(configFile, &conf)
	if err != nil {
		log.Println("Không đọc được file config.")
		panic(err)
	}
	return conf
}

func ConnectDB() *gorm.DB {
	config := SetConfig()
	connectionInfo := fmt.Sprintf(`%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local`, config.Database.User, config.Database.Password, config.Database.Address, config.Database.Database)

	db, err := gorm.Open("mysql", connectionInfo)
	if err != nil {
		panic(err)
	}

	return db
}
