# Build Step
FROM golang:latest AS builder

# Prerequisites
RUN mkdir -p $GOPATH/src/github.com/jplu/visual-search-backend
ADD . $GOPATH/src/github.com/jplu/visual-search-backend
WORKDIR $GOPATH/src/github.com/jplu/visual-search-backend

# Build
ARG build
ARG version
RUN GO111MODULE=on CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o /tmp/visual-search-backend

# Final Step
FROM alpine

# Base packages
RUN apk update && apk upgrade && apk add ca-certificates && update-ca-certificates
RUN apk add --update tzdata
RUN rm -rf /var/cache/apk/*

# Copy binary from build step
COPY --from=builder /tmp/visual-search-backend /home/

# Define timezone
ENV TZ=Europe/Paris

# Define the ENTRYPOINT
WORKDIR /home
ENTRYPOINT ./visual-search-backend

# Document that the service listens on port 8080.
EXPOSE 8080
