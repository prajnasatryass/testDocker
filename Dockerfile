FROM golang:1.23.1-alpine AS build-stage

# Set working directory
WORKDIR /app

# Copy all files into workdir
COPY . .

# Download all dependencies
RUN go mod download

# Build app binary
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /built

# Deploy app binary into a lean image
FROM alpine:3.20 AS release-stage

WORKDIR /

# Copy binary from build-stage to release-stage
COPY --from=build-stage /built /bin

EXPOSE 8080

# Run app binary
ENTRYPOINT ["/bin/built"]