// source: UserManagementService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/database/model"
	"job-seek/pkg/protos"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUserProfile implements CreateUserProfile from UserManagementServiceServer
// generated by protoc-gen-go-grpc.UserManagementServiceServer
func (s UserManagementServiceServerImpl) CreateUserProfile(ctx context.Context, req *protos.UserProfile) (*protos.UserResponse, error) {
	s.log.WithFields(logrus.Fields{
		"request": req,
	}).Info("CreateUserProfile")

	if err := checkUserProfileEmptyFields(req); err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to create user profile, empty fields")
		return nil, err
	}

	acc, err := s.fetchUserAccount(req.UserId, "", "")
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user account")
		return nil, status.Errorf(codes.NotFound, "Failed to fetch user account: %v", err)
	}
	if acc == nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user account, user not found")
		return nil, status.Errorf(codes.NotFound, "Failed to fetch user account: %v", err)
	}

	// Save the user profile to the database
	user, err := s.storeUserProfileToDB(req)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to create user profile")
		return nil, status.Errorf(codes.Internal, "Failed to create user profile: %v", err)
	}

	go func() {
		s.log.WithFields(logrus.Fields{
			"user": user,
		}).Info("User profile created successfully")
		s.log.Info("Start to postprocess user profile")
	}()

	return &protos.UserResponse{
		UserId:  user.Id,
		Status:  "success",
		Message: "User profile created successfully",
	}, nil

	// return nil, status.Errorf(codes.Unimplemented, "method  CreateUserProfile  not implemented")
}

func checkUserProfileEmptyFields(req *protos.UserProfile) error {
	// You can use a database query or any other method to check the email
	// Return true if the email is registered, false otherwise
	if req.Title == "" {
		return status.Errorf(codes.InvalidArgument, "Title cannot be empty")
	}
	if req.UserId == "" {
		return status.Errorf(codes.InvalidArgument, "Reference UserId cannot be empty")
	}
	if req.Position == "" {
		return status.Errorf(codes.InvalidArgument, "Position cannot be empty")
	}
	if req.Description == "" {
		return status.Errorf(codes.InvalidArgument, "Description cannot be empty")
	}

	return nil
}

func (s UserManagementServiceServerImpl) storeUserProfileToDB(user *protos.UserProfile) (*model.UserProfileModel, error) {
	// You can use a database query or any other method to save the user
	// Return an error if the save operation fails, nil otherwise
	instanceModel := &model.UserProfileModel{}
	instanceModel.FromProto(user)
	if err := instanceModel.CreateModel(s.dbClient); err != nil {
		return nil, err
	}
	// Save the user profile to the database
	return instanceModel, nil
}
