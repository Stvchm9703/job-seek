package server_test

import (
	"context"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/services/user_management_service/server"
	"time"

	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
)

var test_run_user_name string = "test_user"
var test_run_user_password string = "test_password"
var test_run_user_id string = ""
var test_run_user_profile_id string = ""

func PretestUserAccountGet(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	testRequest := &protos.GetUserRequest{
		UserName:     &test_run_user_name,
		UserPassword: &test_run_user_password,
	}
	testCtx := context.Background()
	resp, _ := service.GetUserAccount(testCtx, testRequest)

	if resp != nil {
		test_run_user_id = resp.Id
	}
}
func TestUserProfileCreate(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserProfileCreate")
	// Create user profile
	timeStart := time.Now()
	fmt.Println("TestUserProfileCreate")
	testRequest := &protos.UserProfile{
		Id:          "",
		UserId:      test_run_user_id,
		Title:       "test_title",
		Position:    "test_position",
		Description: "test_description",
		Salary:      "20 - 30 k USD",
		Keywords:    []*protos.PreferenceKeyword{},
		Type:        protos.UserProfileType_EMPLOYEE,
		StartDate:   "2021-01-01",
		EndDate:     "2021-12-31",
		Company:     "test_company",
		Location:    "test_location",
	}

	testCtx := context.Background()
	resp, err := service.CreateUserProfile(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "CreateUserProfile",
		}).Error("fail to create user profile from cache")
		fmt.Println("Test CreateUserProfile Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestCreateUserProfile Done")
	pp.Println(" create profile ", resp)
}

func TestUserProfileGet(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserProfileGet")
	// Get user profile
	timeStart := time.Now()
	fmt.Println("TestUserProfileGet")

	// testString := "test_user"
	testRequest := &protos.UserProfile{
		Id: "",
	}

	testCtx := context.Background()
	resp, err := service.GetUserProfile(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "GetUserProfile",
		}).Error("fail to get user profile from cache")
		fmt.Println("Test GetUserProfile Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestGetUserProfile Done")
	pp.Println("resp", resp)
}

func TestUserProfileUpdate(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserProfileUpdate")
	// Update user profile
	timeStart := time.Now()
	fmt.Println("TestUserProfileUpdate")
	testRequest := &protos.UserProfile{
		Id:          test_run_user_id,
		UserId:      test_run_user_id,
		Title:       "test_title",
		Position:    "test_position",
		Description: "test_description",
		Salary:      "20 - 30 k USD",
		Keywords:    []*protos.PreferenceKeyword{},
		Type:        protos.UserProfileType_EMPLOYEE,
		StartDate:   "2021-01-01",
		EndDate:     "2021-12-31",
		Company:     "test_company",
		Location:    "test_location",
	}

	testCtx := context.Background()
	resp, err := service.UpdateUserProfile(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "UpdateUserProfile",
		}).Error("fail to update user profile from cache")
		fmt.Println("Test UpdateUserProfile Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestUpdateUserProfile Done")
}

func TestUserProfileDelete(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserProfileDelete")
	// Delete user profile
	timeStart := time.Now()
	fmt.Println("TestUserProfileDelete")
	testString := "test_user"
	testRequest := &protos.GetUserRequest{
		UserName: &testString,
	}
	testCtx := context.Background()
	resp, err := service.DeleteUserProfile(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "DeleteUserProfile",
		}).Error("fail to delete user profile from cache")
		fmt.Println("Test DeleteUserProfile Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestDeleteUserProfile Done")
}
