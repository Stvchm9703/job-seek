// source: JobSearchService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListCompanyDetail implements ListCompanyDetail from JobSearchServiceServer
// generated by protoc-gen-go-grpc.JobSearchServiceServer
func (s JobSearchServiceServerImpl) ListCompanyDetail(ctx context.Context, req *protos.CompanyDetailRequest) (*protos.CompanyDetailResponse, error) {
	// todo()
	return nil, status.Errorf(codes.Unimplemented, "method  ListCompanyDetail  not implemented")
}
