package sql

import (
	"HHhunter/config"
	"HHhunter/entity"
	"HHhunter/utils"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

func (s *SqlData) Initiate(config *config.ConfigStruct, logger *logrus.Logger) {

	confString := utils.MapToString(config.Repository, logger)

	conf, err := pgxpool.ParseConfig(confString)

	if err != nil {
		logger.Error(err)
	}
	conConf, err1 := pgxpool.ConnectConfig(context.TODO(), conf)

	if err1 != nil {
		logger.Error(err1)
	}

	s.Client = conConf
}

func (s *SqlData) IsExist(singleEntry *entity.ProccessData, logger *logrus.Logger) (isExist bool) {

	connection, err := s.Client.Acquire(context.Background())
	if err != nil {
		logger.Error(err)
	}
	defer connection.Release()

	reply := connection.QueryRow(context.Background(), IsExist, singleEntry.ID)

	//var check interface{}
	err1 := reply.Scan(&isExist)

	if err1 != nil {
		logger.Error(err1)
	}

	return isExist
}
func (s *SqlData) GetSingleData(logger *logrus.Logger) *entity.ProccessData {

	connection, err := s.Client.Acquire(context.Background())
	if err != nil {
		logger.Error(err)
	}
	defer connection.Release()

	reply := connection.QueryRow(context.Background(), SelectEntry, s.Data.ID)
	if err != nil {
		logger.Error(err)
	}

	err1 := reply.Scan(s.Data.ID, s.Data.Name, s.Data.Employer, s.Data.SalaryFrom, s.Data.SalaryTo, s.Data.URL)
	if err1 != nil {
		logger.Error(err1)
	}

	return s.ToProccess(logger)
}
func (s *SqlData) SaveSingleData(singleEntry *entity.ProccessData, logger *logrus.Logger) (isSuccesfull bool) {

	connection, err := s.Client.Acquire(context.Background())
	if err != nil {
		logger.Error(err)
	}
	defer connection.Release()
	s.FromProccess(singleEntry, logger)
	reply, err1 := connection.Exec(context.Background(), InsertEntry, s.Data.ID, s.Data.Name, s.Data.Employer, s.Data.SalaryFrom, s.Data.SalaryTo, s.Data.URL)

	if err1 != nil {
		logger.Error(err1)
	}

	return reply.Insert()
}

func (s *SqlData) ToProccess(logger *logrus.Logger) *entity.ProccessData {

	return &entity.ProccessData{
		ID:         s.Data.ID,
		Name:       s.Data.Name,
		Employer:   s.Data.Employer,
		SalaryFrom: s.Data.SalaryFrom,
		SalaryTo:   s.Data.SalaryTo,
		URL:        s.Data.URL,
	}
}
func (s *SqlData) FromProccess(pds *entity.ProccessData, logger *logrus.Logger) {

	s.Data.ID = pds.ID
	s.Data.Name = pds.Name
	s.Data.Employer = pds.Employer
	s.Data.SalaryFrom = pds.SalaryFrom
	s.Data.SalaryTo = pds.SalaryTo
	s.Data.URL = pds.URL

}
