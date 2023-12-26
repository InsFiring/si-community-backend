package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	rest "si-community/api/v1"
	"si-community/config"

	"github.com/pelletier/go-toml/v2"
)

var (
	configPath string
	tomlConfig config.TomlConfig
)

func init() {
	argumentParser()

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Errorf("파일 읽는 중 에러 발생: ", err)
		return
	}

	err = toml.Unmarshal(content, &tomlConfig)
	if err != nil {
		fmt.Errorf("toml 파일 언마샬링 중 에러 발생: ", err)
		return
	}
}

func argumentParser() {
	flag.StringVar(&configPath, "config", "./config/configuration.toml", "설정파일 풀경로")
	flag.Parse()
}

// @title Your Gin API
// @version 1.0
// @description This is a sample Gin API with Swagger documentation.
// @host localhost:8000
// @BasePath /v1
func main() {
	dbConn, err := config.DBConnection(tomlConfig)
	if err != nil {
		fmt.Errorf("DBConn error: ", err)
		return
	}

	log.Println("Main log...")
	log.Fatal(rest.RunAPI("127.0.0.1:8000", dbConn))
}
