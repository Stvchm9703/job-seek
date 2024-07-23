// source: JobExtendsionService.proto
// Version: 1.0.0

package server
import (
  "context"
  "job-seek/pkg/protos"
  "job-seek/pkg/service_util"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  runConf "job-seek/services/job_extendsion_service/config"
  "sync"
  "net"
  logrus "github.com/sirupsen/logrus"
)


// ApplyJob implements ApplyJob from JobExtendsionServiceServer
// generated by protoc-gen-go-grpc.JobExtendsionServiceServer
func (s JobExtendsionServiceServerImpl) ApplyJob(ctx context.Context, req *protos.JobApply) (*protos.JobResponse, error) {
  // todo()
  return nil, status.Errorf(codes.Unimplemented, "method  ApplyJob  not implemented")
}