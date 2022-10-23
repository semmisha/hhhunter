package hh

import (
	"HHhunter/config"
	"HHhunter/entity"

	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strconv"
)

func (H *HHStruct) Initiate(config *config.ConfigStruct, logger *logrus.Logger) {
	var (
		hhapiaddress = config.Input.ConnectionString
		configMap    = config.Input.RequestConfig
	)

	request, err := http.NewRequest(http.MethodGet, hhapiaddress, nil)
	if err != nil {
		logger.Error(err)
	}
	q := request.URL.Query()
	for i, j := range configMap {
		if len(i) > 0 && len(j) > 0 {
			q.Add(i, j)
		} else {
			logger.Warnln("Empty field: ", i, j)

		}
	}
	request.URL.RawQuery = q.Encode()
	//logger.Infoln(request)

	H.Request = request
}

func (H *HHStruct) GetMultipleData(logger *logrus.Logger) *entity.ProccessDataStruct {

	H.SendRequest(&logrus.Logger{})

	var (
		dataSlice = entity.ProccessDataStruct{}
	)

	for _, i := range H.Items {

		numericID, err := strconv.Atoi(i.Id)
		if err != nil {
			logger.Error(err)
		}
		dataSlice.Items = append(dataSlice.Items, entity.ProccessData{
			ID:         numericID,
			Name:       i.Name,
			Employer:   i.Employer.Name,
			SalaryFrom: i.Salary.From,
			SalaryTo:   i.Salary.To,
			URL:        i.AlternateUrl,
		})
	}

	return &dataSlice
}

func (H *HHStruct) SendRequest(logger *logrus.Logger) {

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	resp, err := client.Do(H.Request)

	if err != nil || resp.StatusCode != http.StatusOK {
		logger.Fatal("Status code %v, error:", resp.StatusCode, err)
	}

	respByte, err := io.ReadAll(resp.Body)

	if err != nil {

		logger.Error(err)

	}

	err = json.Unmarshal(respByte, H)
	if err != nil {
		logger.Error(err)

	}

}
