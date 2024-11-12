# Stage 1: Build the application
FROM golang:1.23 as build

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev 

COPY . /app/

RUN go mod tidy

RUN go build -o wallet-core ./cmd/walletcore

# Stage 2: Run the application
FROM alpine:latest

RUN mkdir -p /opt/app
WORKDIR /opt/app

RUN apk add --no-cache curl bash mysql-client

COPY --from=build /app/wallet-core /opt/app/

COPY internal/database/migrations /opt/app/migrations

COPY entrypoint.sh /opt/app/
RUN chmod +x /opt/app/entrypoint.sh

RUN curl -L -o /usr/bin/migrate https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.deb
RUN chmod +x /usr/bin/migrate

EXPOSE 8080

ENTRYPOINT ["/opt/app/entrypoint.sh"]
