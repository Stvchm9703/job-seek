package server_test

import (
	"context"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/services/job_search_service/server"
	"time"

	"github.com/sirupsen/logrus"
)

func TestJobSearch(service *server.JobSearchServiceServerImpl, log *logrus.Logger) {
	timeStart := time.Now()
	fmt.Println("TestJobSearch")
	minSalary := int32(1000000)
	salaryType := protos.SalaryType_ANNUAL
	classification := int32(6281)
	workLocale := "Sydney"
	allowMixCache := true
	testRequest := &protos.JobSearchRequest{
		UserId:         "32130c50-7afc-4130-9491-4caddc30f81d",
		SalaryType:     &salaryType,
		MinSalary:      &minSalary,
		Keywords:       []string{"golang", "python", "backend"},
		Classification: &classification,
		WorkLocale:     &workLocale,
		AllowMixCache:  &allowMixCache,
	}

	testCtx := context.Background()
	resp, err := service.JobSearch(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "JobSearchServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "JobSearch",
		}).Error("fail to fetch jobs from cache")
		fmt.Println("TestJobSearch Error")
	}
	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("Cooling down for 10 minutes")
	time.Sleep(10 * time.Minute)
	fmt.Println("TestJobSearch Done")
	// pp.Println("TestJobSearch Response", resp)
}
