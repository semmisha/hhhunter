package sql

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type SqlData struct {
	Data struct {
		ID         int
		Name       string
		Employer   string
		SalaryFrom int
		SalaryTo   int
		URL        string
		Added      time.Time
	}
	Client *pgxpool.Pool
}

func NewSqlStruct() *SqlData {

	return &SqlData{}
}
