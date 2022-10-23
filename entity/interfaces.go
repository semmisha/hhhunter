package entity

import (
	"HHhunter/config"

	"github.com/sirupsen/logrus"
)

type Initiator interface {
	Initiate(configStruct *config.ConfigStruct, logger *logrus.Logger)
}
type GetData interface {
	GetMultipleData(logger *logrus.Logger) *ProccessDataStruct
}
type SaveSingleData interface {
	SaveSingleData(singleEntry *ProccessData, logger *logrus.Logger) (isSuccesfull bool)
}

type Config interface {
	GetConfig(configPath string, logger *logrus.Logger)
	SaveConfig(configPath string, logger *logrus.Logger)
}
type InputProcessor interface {
	Initiator
	GetMultipleData(logger *logrus.Logger) *ProccessDataStruct
}
type OutputProcessor interface {
	Initiator
	SaveSingleData
}
type Repository interface {
	Initiator
	SaveSingleData
	GetSingleData(logger *logrus.Logger) *ProccessData
	IsExist(singleEntry *ProccessData, logger *logrus.Logger) (isExist bool)
}
