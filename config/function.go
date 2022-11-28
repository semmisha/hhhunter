package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

func (C *ConfigStruct) GetConfig(filePath string, logger *logrus.Logger) {

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		logger.Error(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {

		logger.Error(err)

	}
	err1 := yaml.Unmarshal(data, C)
	if err1 != nil {
		logger.Fatal(err1)

	}
	if C == nil {
		logger.Fatal("Config is empty!!")
	}
	//logger.Infoln(C)

}

func (C *ConfigStruct) SaveConfig(filePath string, logger *logrus.Logger) {}
