FROM golang:1.15-alpine

RUN apk add --no-cache git

# Add docker compose wait for database https://github.com/ufoscout/docker-compose-wait
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

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
RUN go build -o ./bin/GoGinApi ./cmd/server


# This container exposes port 8082 to the outside world
EXPOSE 8082

# Run the binary program produced by `go install`
CMD ["./bin/GoGinApi"]