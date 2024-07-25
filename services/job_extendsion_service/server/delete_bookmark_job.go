// source: JobExtendsionService.proto
// Version: 1.0.0

package server

import (
	"context"
	"job-seek/pkg/protos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteJobBookmark implements DeleteJobBookmark from JobExtendsionServiceServer
// generated by protoc-gen-go-grpc.JobExtendsionServiceServer
func (s JobExtendsionServiceServerImpl) DeleteJobBookmark(ctx context.Context, req *protos.JobBookmarkRequest) (*protos.JobBookmarkResponse, error) {
	// todo()
	return nil, status.Errorf(codes.Unimplemented, "method  DeleteJobBookmark  not implemented")
}
