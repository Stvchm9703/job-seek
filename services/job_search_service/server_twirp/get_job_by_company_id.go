// source: JobSearchService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/protos"

	"github.com/twitchtv/twirp"
)

// GetJobByCompanyId implements GetJobByCompanyId from JobSearchServiceServer
// generated by protoc-gen-go-grpc.JobSearchServiceServer
func (s *JobSearchServiceServerImpl) GetJobByCompanyId(ctx context.Context, req *protos.JobSearchRequest) (*protos.JobSearchResponse, error) {
	// todo()
	// return nil, status.Errorf(codes.Unimplemented, "method  GetJobByCompanyId  not implemented")
	return nil, twirp.Unauthenticated.Error("method  GetJobByCompanyId  not implemented")
}
