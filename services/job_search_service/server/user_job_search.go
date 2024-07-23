// source: JobSearchService.proto
// Version: 1.0.0

package server
import (
  "context"
  "job-seek/pkg/protos"
  "job-seek/pkg/service_util"
  "google.golang.org/grpc"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  runConf "job-seek/services/job_search_service/config"
  "sync"
  "net"
  logrus "github.com/sirupsen/logrus"
)


// UserJobSearch implements UserJobSearch from JobSearchServiceServer
// generated by protoc-gen-go-grpc.JobSearchServiceServer
func (s JobSearchServiceServerImpl) UserJobSearch(ctx context.Context, req *protos.JobSearchRequest) (*protos.JobSearchResponse, error) {
  // todo()
  return nil, status.Errorf(codes.Unimplemented, "method  UserJobSearch  not implemented")
}