generate-proto:
	protoc \
	--proto_path=pkg/protos \
	--go_out=pkg/protos --go_opt=paths=source_relative \
	--go-grpc_out=pkg/protos --go-grpc_opt=paths=source_relative \
		pkg/protos/*.proto

sync-proto:
	cp -r pkg/protos/* ~/git_src/job-seek/

build-fetch-job-service:
	go build -o bin/fetch-job-service services/fetch_job_service/main.go


build-service:\
 build-fetch-job-service
