package main

import (
	"HHhunter/config"
	"HHhunter/controller/api/hh"
	"HHhunter/controller/api/telegram"
	"HHhunter/entity"
	"HHhunter/logging"
	"HHhunter/repository/sql"
	"fmt"
	"time"
)

const (
	ConfigPath = "/app/config/files/config.yaml"
)

func main() {

	logger := logging.Logger()

	conf := config.NewConfigStruct()
	conf.GetConfig(ConfigPath, logger)

	pi := entity.ProccessorInterfaces{
		hh.NewHHStruct(),
		telegram.NewTelegramStruct(),
		sql.NewSqlStruct(),
	}
	pi.Initiate(conf, logger)
	for {
		resultStatistic := pi.Process(logger)
		fmt.Printf("\n%+v\n", resultStatistic)
		time.Sleep(3 * time.Hour)
	}
}
