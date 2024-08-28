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

	log.Info("Check UserCvProfileModel")
	ucvm := &model.UserCVProfileModel{}
	err = ucvm.DefineModel(dbClient)
	if err != nil {
		log.WithFields(logrus.Fields{
			"config": runConf.RuntimeConfig,
			"model":  "UserCVProfileModel",
			"error":  err,
		}).Fatal("Failed to define")
	}

	log.Info("completed")

}
