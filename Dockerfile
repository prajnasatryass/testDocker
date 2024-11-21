FROM golang:1.23.1-alpine AS build-stage

# Set working directory
WORKDIR /app

# Copy all files into workdir
COPY . .

# Download all dependencies
RUN go mod download

## Install dependencies and add wait-for-it script
#RUN apk add --no-cache bash curl
#ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
#RUN chmod +x /wait-for-it.sh

## Wait for the database to be ready, then run app to verify DB is reachable
#RUN /wait-for-it.sh db:5432 --timeout=15 --strict -- go run -mod=mod ./main.go

# Wait for the database to be ready, then run app to verify DB is reachable
RUN go run -mod=mod ./main.go