// source: UserManagementService.proto
// Version: 1.0.0

package server

import (
	"context"
	"fmt"
	"job-seek/pkg/database/model"
	"job-seek/pkg/protos"

	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateUserProfile implements UpdateUserProfile from UserManagementServiceServer
// generated by protoc-gen-go-grpc.UserManagementServiceServer
func (s UserManagementServiceServerImpl) UpdateUserProfile(ctx context.Context, req *protos.UserProfile) (*protos.UserResponse, error) {

	s.log.WithFields(logrus.Fields{
		"request": req,
	}).Info("UpdateUserProfile")

	if req.Id == "" {
		s.log.WithFields(logrus.Fields{
			"request": req,
		}).Error("Failed to update user profile, empty fields")
		return nil, status.Errorf(codes.InvalidArgument, "Failed to update user profile")
	}

	// Update the user profile
	result, err := s.updateUserProfile(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user profile: %v", err)
	}

	return result, nil
}

func (s UserManagementServiceServerImpl) updateUserProfile(req *protos.UserProfile) (*protos.UserResponse, error) {
	// You can use a database query or any other method to update the user
	// Return an error if the update operation fails, nil otherwise
	instanceModel := &model.UserProfileModel{}
	instanceModel.FromProto(req)

	if err := instanceModel.UpdateModel(s.dbClient); err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to update user profile")
		return nil, err
	}

	return &protos.UserResponse{
		UserId:  fmt.Sprintf("%d", instanceModel.ID),
		Message: "User profile updated successfully",
		Status:  "success",
	}, nil
}
