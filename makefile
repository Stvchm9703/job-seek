current_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
rust_dir := $(current_dir)../job-seek-v2

generate-go-proto:
	protoc \
	--proto_path=pkg/protos \
	--go_out=pkg/protos --go_opt=paths=source_relative \
	--go-grpc_out=pkg/protos --go-grpc_opt=paths=source_relative \
		pkg/protos/*.proto

sync-proto:
	cp -r pkg/protos/* $(rust_dir)/packages/protos/defines/

check-rust-protos:
	cd $(rust_dir)/packages/protos && \
	cargo check &&\
	cd $(current_dir)

generate-proto: \
	generate-go-proto \
	sync-proto \
	check-rust-protos

build-job-search-service:
	go build -o bin/job-search-service services/job-search-service/main.go


build-service:\
	build-job-search-service

tool-generate-service-template:
	go run tools/generate-service-template/main.go