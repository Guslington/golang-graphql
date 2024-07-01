# build stage
FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /go/src/app

# Copy local code to the container image.
# .dockerignore is used here to exclude unwanted files from the copy
COPY go.* .
COPY pkg/ /go/src/app/pkg/
COPY services/ /go/src/app/services/

# Build the binary.
RUN go mod download
RUN go generate services/movies/...
RUN go build -o /go/bin/app services/movies/server.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# copy in the go binary from the build stage
COPY --from=builder /go/bin/app /app

# Run the api service on container startup.
ENTRYPOINT /app

# Expose the api port to allow external requests
EXPOSE 8080