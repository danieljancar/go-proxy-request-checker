package reportgenerator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Report struct {
	Name     string
	Date     string
	Status   string
	Total    int
	Success  int
	Fail     int
	Requests []RequestReport
}

type RequestReport struct {
	URL              string
	ExpectedResponse int
	ActualResponse   int
}

func NewReport() *Report {
	return &Report{
		Name:     "Report",
		Date:     time.Now().Format("2006-01-02 15:04:05"),
		Status:   "Success",
		Total:    0,
		Success:  0,
		Fail:     0,
		Requests: []RequestReport{},
	}
}

func (r *Report) Analyze() {
	for _, request := range r.Requests {
		r.Total++
		if request.ExpectedResponse == request.ActualResponse {
			r.Success++
		} else {
			r.Fail++
			r.Status = "Fail"
		}
	}
}

func (r *Report) AddRequest(url string, expectedResponse, actualResponse int) {
	requestReport := RequestReport{
		URL:              url,
		ExpectedResponse: expectedResponse,
		ActualResponse:   actualResponse,
	}
	r.Requests = append(r.Requests, requestReport)
}

func (r *Report) SaveToFile(filePath string) error {
	r.Analyze()
	fileData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		log.Printf("Error transfering report to JSON: %s", err)
		return err
	}

	err = ioutil.WriteFile(filePath, fileData, 0644)
	if err != nil {
		log.Printf("Error writing report to file: %s", err)
		return err
	}

	log.Printf("Report successfully saved to %s", filePath)
	return nil
}
