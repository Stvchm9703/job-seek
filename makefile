current_dir := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))
rust_dir := $(current_dir)../job-seek-v2
python_dir := $(current_dir)../job-seek-v3
node_dir := $(current_dir)../job-seek-web-app
generate-go-proto:
	protoc \
	--proto_path=pkg/protos \
	--go_out=pkg/protos --go_opt=paths=source_relative \
	--go-grpc_out=pkg/protos --go-grpc_opt=paths=source_relative \
		pkg/protos/*.proto

sync-proto:
	# cp -r pkg/protos/*.proto $(rust_dir)/packages/protos/defines/
	cp -r pkg/protos/*.proto $(python_dir)/app/libs/protos/
	cp -r pkg/protos/*.proto $(node_dir)/protos/

check-rust-protos:
	cd $(rust_dir)/packages/protos && \
	cargo check &&\
	cd $(current_dir)

generate-proto: \
	generate-go-proto \
	sync-proto \
	# check-rust-protos

build-job-search-service:
	go build -o bin/job-search-service services/job_search_service/main.go

run-job-search-service:
	go run services/job_search_service/main.go  -C=config/job_search_service.toml run -V=4


build-service:\
	build-job-search-service

tool-generate-service-template:
	go run tools/generate-service-template/main.go


# db table init 
init-db-job_search_service:
	go run ./services/job_search_service/main.go db init -C=./config/job_search_service.toml  

init-db-user_service:
	go run ././services/user_management_service/main.go db init -C=./config/user_account_service.toml

init-db:\
	init-db-job_search_service \
	init-db-user_service