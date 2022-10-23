package entity

type ProccessorInterfaces struct {
	Input      InputProcessor
	Output     OutputProcessor
	Repository Repository
}

type ProccessDataStruct struct {
	Items []ProccessData
}

type ProccessData struct {
	ID         int
	Name       string
	Employer   string
	SalaryFrom int
	SalaryTo   int
	URL        string
}

type Count struct {
	Success      uint32
	FailedSend   uint32
	FailedSaved  uint32
	AlreadyExist uint32
}

func NewCount() *Count {
	return &Count{Success: 0, FailedSaved: 0, FailedSend: 0, AlreadyExist: 0}
}
