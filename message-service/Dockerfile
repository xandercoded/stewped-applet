FROM golang:1.15-alpine AS build

COPY . /go/src/app
WORKDIR /go/src/app
RUN go mod download
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/message-service /go/src/app/message-service

FROM scratch
COPY --from=build /go/bin/message-service /go/bin/
EXPOSE 8080
ENTRYPOINT ["/go/bin/message-service"]
