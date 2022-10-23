package telegram

import (
	"HHhunter/config"
	"HHhunter/entity"
	"fmt"
	"github.com/semmisha/ClientAPI"
	"github.com/sirupsen/logrus"
)

func (T *TelegramStruct) Initiate(config *config.ConfigStruct, logger *logrus.Logger) {
	T.Client = ClientAPI.NewConnection(config.Output["channel"], config.Output["address"])

}
func (T *TelegramStruct) SaveSingleData(singleEntry *entity.ProccessData, logger *logrus.Logger) (isSuccesfull bool) {

	//buf := bytes.NewBufferString(fmt.Sprint(singleEntry))
	if singleEntry == nil {

		logger.Error("empty entry ProccessData", singleEntry)
		return false
	}

	_, err := T.Client.Write([]byte(SuitUp(singleEntry, logger)))
	if err != nil {
		logger.Error(err)
		return false
	}
	return true
}

func SuitUp(singleEntry *entity.ProccessData, logger *logrus.Logger) string {

	return fmt.Sprintf(" ID:%v\n Name:%v\n Employer:%v\n SalaryFrom:%v\n SaleryTo:%v\n, URL:%v\n", singleEntry.ID, singleEntry.Name, singleEntry.Employer, singleEntry.SalaryFrom, singleEntry.SalaryTo, singleEntry.URL)

}
