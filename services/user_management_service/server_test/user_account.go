package server_test

import (
	"context"
	"fmt"
	"job-seek/pkg/protos"
	"job-seek/services/user_management_service/server"
	"time"

	"github.com/sirupsen/logrus"
)

var test_user_account_id string = ""

// / TestUserAccountCreate is a test function to create a user account
func TestUserAccountCreate(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserAccountCreate")
	// Create user account
	timeStart := time.Now()
	fmt.Println("TestUserAccountCreate")
	testRequest := &protos.UserAccount{
		Id:           "",
		UserName:     "test_user",
		UserPassword: "test_password",
		UserEmail:    "test_user@mail.com",
		UserAddress:  "test_address",
		UserPhone:    "1234567890",
	}
	testCtx := context.Background()
	resp, err := service.CreateUserAccount(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "CreateUserAccount",
		}).Error("fail to create user account from cache")
		fmt.Println("Test CreateUserAccount Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestCreateUserAccount Done")
	// pp.Println(resp)
	test_user_account_id = resp.GetUserId()
}

// / TestUserAccountGet is a test function for GetUserAccount
func TestUserAccountGet(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserAccountGet")
	// Get user account
	timeStart := time.Now()
	fmt.Println("TestUserAccountGet")
	testString := "test_user"
	// testPass := "test_password"
	testRequest := &protos.GetUserRequest{
		UserName: &testString,
	}
	testCtx := context.Background()
	resp, err := service.GetUserAccount(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "GetUserAccount",
		}).Error("fail to get user account from cache")
		fmt.Println("Test GetUserAccount Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestCreateUserAccount Done")

}

// / TestUserAccountUpdate is a test function for UpdateUserAccount
func TestUserAccountUpdate(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
	log.Info("Start TestUserAccountUpdate")
	// Update user account
	timeStart := time.Now()
	fmt.Println("TestUserAccountUpdate")
	testRequest := &protos.UserAccount{
		Id:       test_user_account_id,
		UserName: "test_user",
		// UserPassword: "test_password",
		UserEmail:   "test_user@mail.com",
		UserAddress: "test_address",
		UserPhone:   "1234567890",
	}
	testCtx := context.Background()
	resp, err := service.UpdateUserAccount(testCtx, testRequest)
	if err != nil {
		log.WithFields(map[string]interface{}{
			"model":    "UserManagementServiceServerImpl",
			"error":    err,
			"request":  testRequest,
			"response": resp,
			"method":   "UpdateUserAccount",
		}).Error("fail to update user account from cache")
		fmt.Println("TestUserAccountUpdate Error")
	}

	timeTook := time.Since(timeStart)
	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
	fmt.Println("TestUpdateUserAccount Done")
}

// // / TestUserAccountDelete is a test function for DeleteUserAccount
// func TestUserAccountDelete(service *server.UserManagementServiceServerImpl, log *logrus.Logger) {
// 	log.Info("Start TestUserAccountDelete")
// 	// Delete user account
// 	timeStart := time.Now()
// 	fmt.Println("TestUserAccountDelete")
// 	testString := "test_user"
// 	testRequest := &protos.GetUserRequest{
// 		UserName: &testString,
// 	}
// 	testCtx := context.Background()
// 	resp, err := service.DeleteUserProfile(testCtx, testRequest)
// 	if err != nil {
// 		log.WithFields(map[string]interface{}{
// 			"model":    "UserManagementServiceServerImpl",
// 			"error":    err,
// 			"request":  testRequest,
// 			"response": resp,
// 			"method":   "DeleteUserProfile",
// 		}).Error("fail to delete user account from cache")
// 		fmt.Println("TestDeleteUserProfile Error")
// 	}

// 	timeTook := time.Since(timeStart)
// 	fmt.Printf("Time took : %f secs \n", timeTook.Seconds())
// 	fmt.Println("TestDeleteUserProfile Done")
// }
