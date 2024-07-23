// source: UserManagementService.proto
// Version: 1.0.0

package server
import (
  "context"
  "job-seek/pkg/protos"
  "job-seek/pkg/service_util"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  runConf "job-seek/services/user_management_service/config"
  "sync"
  "net"
  logrus "github.com/sirupsen/logrus"
)


// GetUserAccount implements GetUserAccount from UserManagementServiceServer
// generated by protoc-gen-go-grpc.UserManagementServiceServer
func (s UserManagementServiceServerImpl) GetUserAccount(ctx context.Context, req *protos.GetUserRequest) (*protos.UserAccount, error) {
  // todo()
  return nil, status.Errorf(codes.Unimplemented, "method  GetUserAccount  not implemented")
}