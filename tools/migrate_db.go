package main

import (
	config "job-seek/pkg/config"
	"job-seek/pkg/database"
	"job-seek/pkg/database/model"
	serviceConf "job-seek/services/job_search_service/config"

	logrus "github.com/sirupsen/logrus"
)

// migration script to migrate the database

func main() {
	// connect to the database
	rtConfig := serviceConf.ServiceConfig{
		Host:               "localhost",
		Port:               60020,
		Server:             config.ServerConfig{},
		MeiliSearchService: config.DatabaseConfig{},
		SurrealDBService:   config.DatabaseConfig{},
		DBService: config.DatabaseConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "job_search_service",
			Password: "service_job_search",
		},
	}

	dbClient, err := database.InitConnection(&rtConfig.DBService, "development")
	log := logrus.New()

	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
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
			"config": rtConfig,
			"model":  "CompanyDetailModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check JobModel")
	jobPost := &model.JobModel{}
	err = jobPost.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "JobModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check JobCacheModel")
	jobCache := &model.JobCacheListModel{}
	err = jobCache.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "JobCacheListModel",
			"error":  err,
		}).Fatal("Failed to connect to database")
	}

	log.Info("Check UserAccountModel")
	uam := &model.UserAccountModel{}
	err = uam.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "UserAccountModel",
			"error":  err,
		}).Fatal("Failed to define")
	}
	log.Info("Check PreferenceKeywordModel")
	pkm := &model.PreferenceKeywordModel{}
	err = pkm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "PreferenceKeywordModel",
			"error":  err,
		}).Fatal("Failed to define")
	}
	log.Info("Check UserProfileModel")
	upm := &model.UserProfileModel{}
	err = upm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "UserProfileModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyUserPreferenceModel")
	supm := &model.SurveyUserPreferenceModel{}
	err = supm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "SurveyUserPreferenceModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyJobQuestionModel")
	sqpm := &model.SurveyJobQuestionModel{}
	err = sqpm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "SurveyJobQuestionModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyJobPreferenceModel")
	sjpm := &model.SurveyJobPreferenceModel{}
	err = sjpm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": rtConfig,
			"model":  "SurveyJobPreferenceModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("completed")
}
