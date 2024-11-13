# Stage 1: Build the application
FROM golang:1.23 as build

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev 

COPY . /app/

RUN go mod tidy

RUN go build -o wallet-core ./cmd/walletcore

# Stage 2: Run the application
FROM golang:1.23

RUN mkdir -p /opt/app
WORKDIR /opt/app

RUN apt-get update && apt-get install -y curl bash default-mysql-client dpkg

COPY --from=build /app/wallet-core /opt/app/

COPY internal/database/migrations /opt/app/migrations

COPY entrypoint.sh /opt/app/
RUN chmod +x /opt/app/entrypoint.sh

RUN curl -L -O https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.deb
RUN dpkg -i migrate.linux-amd64.deb  
RUN rm -f migrate.linux-amd64.deb

EXPOSE 8080

ENTRYPOINT ["/opt/app/entrypoint.sh"]
