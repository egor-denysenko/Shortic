FROM golang:1.19.0-alpine3.16 as builder

WORKDIR /go/src
RUN mkdir build
COPY . .
RUN mv .env ./build/
# Required build arguments
ARG version
ARG commitHash
ARG buildUser
ARG buildTime

#Debug build arguments
#RUN echo $version
#RUN echo $commitHash
#RUN echo $buildUser
#RUN echo $buildTime

# Compiling stage
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'shortic/pkg/deployment.version=$version' -X 'shortic/pkg/deployment.commitHash=$commitHash' -X 'shortic/pkg/deployment.buildUser=$buildUser' -X 'shortic/pkg/deployment.buildTime=$buildTime'" -o ./build/shortic ./cmd/restapi/main.go


FROM alpine:3.16.2
WORKDIR /cmd
# Retrieving the compiled binary from the stage before
COPY --from=builder /go/src/build .
CMD ["./shortic"]
