// source: JobExtendsionService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateJobApply implements UpdateJobApply from JobExtendsionServiceServer
// generated by protoc-gen-go-grpc.JobExtendsionServiceServer
func (s JobExtendsionServiceServerImpl) UpdateJobApply(ctx context.Context, req *protos.JobApply) (*protos.JobResponse, error) {
	// todo()
	return nil, status.Errorf(codes.Unimplemented, "method  UpdateJobApply  not implemented")
}
