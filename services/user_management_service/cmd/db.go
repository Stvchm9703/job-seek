package cmd

import (
	"job-seek/pkg/database"
	"job-seek/pkg/database/model"
	runConf "job-seek/services/user_management_service/config"

	logrus "github.com/sirupsen/logrus"
)

func InitDB(verboseLevel int) {
	initLogger(3, verboseLevel)
	log.Info("Start Database Check")
	// ctx, ctxCancel := context.WithCancel(context.Background())
	// defer ctxCancel()
	dbClient, err := database.InitConnection(&runConf.RuntimeConfig.DBService, "development")
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"error":  err,
		}).Fatal("Failed to connect to database")
	}
	log.Info("Connected to database")
	log.Info("Check database model")
	// Check database model
	log.Info("Check UserAccountModel")
	uam := &model.UserAccountModel{}
	err = uam.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "UserAccountModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check UserProfileModel")
	upm := &model.UserProfileModel{}
	err = upm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "UserProfileModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check PreferenceKeywordModel")
	pkm := &model.PreferenceKeywordModel{}
	err = pkm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "PreferenceKeywordModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyUserPreferenceModel")
	supm := &model.SurveyUserPreferenceModel{}
	err = supm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "SurveyUserPreferenceModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyJobQuestionModel")
	sqpm := &model.SurveyJobQuestionModel{}
	err = sqpm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "SurveyJobQuestionModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("Check SurveyJobPreferenceModel")
	sjpm := &model.SurveyJobPreferenceModel{}
	err = sjpm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "SurveyJobPreferenceModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	// log.Info("Check UserCvProfileModel")
	// ucvm := &model.UserCVProfileModel{}
	// err = ucvm.DefineModel(dbClient)
	// if err != nil {
	// 	log.WithFields(logrus.Fields{
	// 		"config": runConf.RuntimeConfig,
	// 		"model":  "UserCVProfileModel",
	// 		"error":  err,
	// 	}).Fatal("Failed to define")
	// }

	log.Info("completed")

}
