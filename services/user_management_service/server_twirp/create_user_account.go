// source: UserManagementService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/database_v1/model"
	"job-seek/pkg/protos"

	"github.com/sirupsen/logrus"
	"github.com/twitchtv/twirp"
)

// CreateUserAccount implements CreateUserAccount from UserManagementServiceServer
// generated by protoc-gen-go-grpc.UserManagementServiceServer
func (s *UserManagementServiceServerImpl) CreateUserAccount(ctx context.Context, req *protos.UserAccount) (*protos.UserResponse, error) {

	// Create the user account
	if err := checkUserAccountEmptyFields(req); err != nil {
		return nil, err
	}

	// Save the user account to the database
	user, err := s.storeUserAccountToDB(req)
	if err != nil {
		// return nil, status.Errorf(codes.Internal, "Failed to create user account: %v", err)
		return nil, twirp.InternalErrorWith(err)
	}

	s.log.WithFields(logrus.Fields{
		"user": user,
	}).Info("User account created successfully")

	// Return the user response
	return &protos.UserResponse{
		// UserId:  fmt.Sprintf("%d", user.ID),
		UserId:  user.ID,
		Status:  "success",
		Message: "User account created successfully",
	}, nil
}

func checkUserAccountEmptyFields(req *protos.UserAccount) error {
	// You can use a database query or any other method to check the email
	// Return true if the email is registered, false otherwise
	if req.UserEmail == "" {
		// return status.Errorf(codes.InvalidArgument, "Email cannot be empty")
		return twirp.InvalidArgumentError("UserEmail", "Email cannot be empty")
	}
	// if req.UserPassword == "" {
	// 	return status.Errorf(codes.InvalidArgument, "Password cannot be empty")
	// }
	if req.UserName == "" {
		// return status.Errorf(codes.InvalidArgument, "Name cannot be empty")
		return twirp.InvalidArgumentError("UserName", "Name cannot be empty")
	}
	return nil
}

func (s *UserManagementServiceServerImpl) storeUserAccountToDB(user *protos.UserAccount) (*model.UserAccountModel, error) {
	// You can use a database query or any other method to save the user
	// Return an error if the save operation fails, nil otherwise
	instanceModel := &model.UserAccountModel{}
	instanceModel.FromProto(user)

	// Save the user account to the database
	// You can use a database query or any other method to save the user
	// Return an error if the save operation fails, nil otherwise
	if err := instanceModel.CreateModel(s.dbClient); err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
			"user":  user,
		}).Error("Failed to save user account to database")
		return nil, err
	}

	result, err := instanceModel.GetModelByWildKey(s.dbClient)

	if err != nil {
		s.log.WithFields(logrus.Fields{
			"error": err,
			"user":  user,
		}).Error("Failed to get user account from database")
		return instanceModel, err
	}
	instanceModel.FromProto(result)

	return instanceModel, nil
}
