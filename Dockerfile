#FROM golang:1.14-alpine
#
#RUN apk add --no-cache git
#
## ENV GO111MODULE=on
#
## Set the Current Working Directory inside the container
#WORKDIR /app/GoGinApi
#
## We want to populate the module cache based on the go.{mod,sum} files.
#COPY go.mod .
#COPY go.sum .
#COPY .env .
#
#RUN go mod download
#
#COPY . .
#
##RUN CGO_ENABLED=0 go test -v
#
## Build the Go app
#RUN go build -o ./bin/GoGinApi ./cmd/server
#
#
## This container exposes port 8080 to the outside world
#EXPOSE 8082
#
## Run the binary program produced by `go install`
#CMD ["./out/GoGinApi"]

###################################################################################################


# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.14 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go modules.
# Allows container builds to reuse downloaded dependencies.
COPY go.* ./
COPY go.mod ./
COPY go.sum ./

RUN go mod download

# Copy local code to the container image.
COPY . ./

#RUN CGO_ENABLED=0 go test -v

# Build the binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o ./bin/GoGinApi server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

EXPOSE 8082

# Run the web service on container startup.
CMD ["/server"]