package cmd

import (
	"job-seek/pkg/database"
	"job-seek/pkg/database/model"
	runConf "job-seek/services/job_search_service/config"

	logrus "github.com/sirupsen/logrus"
)

func InitDB(verboseLevel int) {
	initLogger(3, verboseLevel)
	log.Info("Start Database Check")
	// ctx, ctxCancel := context.WithCancel(context.Background())
	// defer ctxCancel()
	dbClient, err := database.InitConnection(&runConf.RuntimeConfig.SurrealDBService, "development")
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"error":  err,
		}).Fatal("Failed to connect to database")
	}
	log.Info("Connected to database")
	log.Info("Check database model")
	// Check database model
	log.Info("Check CompanyDetailModel")
	companyDetail := &model.CompanyDetailModel{}
	err = companyDetail.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "CompanyDetailModel",
			"error":  err,
		}).Fatal("Failed to connect to database")
	}

	log.Info("Check JobModel")
	jobPost := &model.JobModel{}
	err = jobPost.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "JobModel",
			"error":  err,
		}).Fatal("Failed to connect to database")
	}

	log.Info("Check JobCacheModel")
	jobCache := &model.JobCacheListModel{}
	err = jobCache.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "JobCacheListModel",
			"error":  err,
		}).Fatal("Failed to connect to database")
	}

}
