package entity

import (
	"HHhunter/config"
	"github.com/sirupsen/logrus"
)

func (P *ProccessorInterfaces) Process(logger *logrus.Logger) *Count {

	var (
		input      = P.Input
		output     = P.Output
		repository = P.Repository
	)

	processData := input.GetMultipleData(logger)
	count := NewCount()
	for _, i := range processData.Items {

		if ok := repository.IsExist(&i, logger); !ok {

			if ok := output.SaveSingleData(&i, logger); !ok {
				logger.Error("Unable to send to Output", i.Name, i.ID)
				count.FailedSend++
				continue

			}
			if ok := repository.SaveSingleData(&i, logger); !ok {
				logger.Warn("Unable to save in Repository", i.Name, i.ID)
				count.FailedSaved++
				continue
			}
			count.Success++
		}
		count.AlreadyExist++
	}

	return count
}

func (P *ProccessorInterfaces) Initiate(config *config.ConfigStruct, logger *logrus.Logger) {

	P.Input.Initiate(config, logger)
	P.Output.Initiate(config, logger)
	P.Repository.Initiate(config, logger)

}
