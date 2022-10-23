package sql

const (
	IsExist     = "select exists (select id from hhhunter where id=$1)"
	SelectEntry = "select id from hhhunter where id=$1"
	SelectAll   = "select (id, name, employer,salaryfrom, salaryto, url, added) from hhhunter"
	InsertEntry = "insert into hhhunter (\"id\", name, employer, salaryfrom,salaryto,url) values ($1,$2,$3,$4,$5,$6)"
	UpdateEntry = ""
)
