// source: UserManagementService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/database/model"
	"job-seek/pkg/protos"

	logrus "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUserProfile implements GetUserProfile from UserManagementServiceServer
// generated by protoc-gen-go-grpc.UserManagementServiceServer
func (s UserManagementServiceServerImpl) ListUserProfile(ctx context.Context, req *protos.GetUserRequest) (*protos.ListUserProfileResponse, error) {
	s.log.WithFields(logrus.Fields{
		"request": req,
	}).Info("GetUserProfile")

	if err := checkFetchUserAccountEmptyFields(req); err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user profile, empty fields")
		return nil, err
	}

	profiles, err := s.fetchUserProfiles(req.GetUserId())
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user profile")
		return nil, status.Errorf(codes.Internal, "Failed to fetch user profile: %v", err)
	}
	if profiles == nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user profile, user not found")
		return nil, status.Errorf(codes.NotFound, "Failed to fetch user profile: %v", err)
	}

	return &protos.ListUserProfileResponse{
		Status:   "success",
		Message:  "User profile fetched successfully",
		UserId:   req.GetUserId(),
		Profiles: profiles,
	}, nil
	// return nil, status.Errorf(codes.Unimplemented, "method  GetUserProfile  not implemented")
}

func (s UserManagementServiceServerImpl) fetchUserProfiles(userId string) ([]*protos.UserProfile, error) {
	// You can use a database query or any other method to save the user
	// Return an error if the save operation fails, nil otherwise
	instanceModel := &model.UserProfileModel{
		UserId: userId,
	}
	user, err := instanceModel.GetModelByUserId(s.dbClient)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to fetch user profile")
		return nil, err
	}
	return user, nil
}