package hh

import "net/http"

type HHStruct struct {
	Request *http.Request
	Items   []struct {
		Id         string      `json:"id"`
		Name       string      `json:"name"`
		Department interface{} `json:"department"`
		//HasTest                bool        `json:"has_test"`
		//ResponseLetterRequired bool        `json:"response_letter_required"`
		Area struct {
			//	Id   string `json:"id"`
			Name string `json:"name"`
			//	Url  string `json:"url"`
		} `json:"area"`
		Salary struct {
			From int `json:"from"`
			To   int `json:"to"`
			//	Currency string `json:"currency"`
			Gross bool `json:"gross"`
		} `json:"salary"`
		Type struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		//Address           interface{}   `json:"address"`
		ResponseUrl interface{} `json:"response_url"`
		//SortPointDistance interface{}   `json:"sort_point_distance"`
		PublishedAt string `json:"published_at"`
		CreatedAt   string `json:"created_at"`
		//Archived          bool          `json:"archived"`
		//ApplyAlternateUrl string      `json:"apply_alternate_url"`
		//InsiderInterview interface{} `json:"insider_interview"`
		//Url              string      `json:"url"`
		//AdvResponseUrl    string        `json:"adv_response_url"`
		AlternateUrl string `json:"alternate_url"`
		//Relations    []interface{} `json:"relations"`
		Employer struct {
			//	Id           string `json:"id"`
			Name string `json:"name"`
			//	Url          string `json:"url"`
			//	AlternateUrl string `json:"alternate_url"`
			//	LogoUrls     struct {
			//		Field1   string `json:"90"`
			//		Field2   string `json:"240"`
			//		Original string `json:"original"`
			//	} `json:"logo_urls"`
			//VacanciesUrl string `json:"vacancies_url"`
			//Trusted      bool   `json:"trusted"`
		} `json:"employer"`
		//Snippet struct {
		//	Requirement    string `json:"requirement"`
		//	Responsibility string `json:"responsibility"`
		//} `json:"snippet"`
		//Contacts interface{} `json:"contacts"`
		//Schedule struct {
		//	Id   string `json:"id"`
		//	Name string `json:"name"`
		//} `json:"schedule"`
		//WorkingDays          []interface{} `json:"working_days"`
		//WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
		//WorkingTimeModes     []interface{} `json:"working_time_modes"`
		//AcceptTemporary      bool          `json:"accept_temporary"`
	} `json:"items"`
}

func NewHHStruct() *HHStruct {
	return &HHStruct{}
}
