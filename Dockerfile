FROM golang:1.14-alpine

RUN apk add --no-cache git

# ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app/GoGinApi

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
COPY .env .

RUN go mod download

COPY . .

#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/GoGinApi .


# This container exposes port 8080 to the outside world
EXPOSE 8082

# Run the binary program produced by `go install`
CMD ["./out/GoGinApi"]