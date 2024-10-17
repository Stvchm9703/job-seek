# build process for job search service
FROM    golang:1.23.2 AS build-stage
WORKDIR /app

COPY    go.mod go.sum 
RUN     go mod download

COPY    services/user_management_service/ ./services/user_management_service/
COPY    pkg/ ./pkg/
COPY    config/user_management_service.toml ./base.toml

RUN     go build -o service_app ./services/user_management_service/main.go

# # test stage
# FROM    build-stage AS test-stage
# RUN     go test -v ./...



# final stage
FROM    alpine:3.10 AS production-stage
WORKDIR /

COPY    --from=build-stage /app/service_app .
COPY    --from=build-stage /app/base.toml .

EXPOSE  60020

USER   nonroot:nonroot
CMD    ["./service_app" , "-C", "base.toml" ]
