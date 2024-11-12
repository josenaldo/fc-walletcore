#!/bin/bash

# Wait until MySQL is ready
echo "Waiting for MySQL to be ready..."

until mysqladmin ping -h "mysql-wallet" --silent; do
  echo "Waiting for mysql-wallet..."
  sleep 2
done

echo "MySQL is up and running."

# Run database migrations
migrate -path=/app/migrations -database "mysql://root:root@tcp(mysql-wallet:3306)/wallet" -verbose up

# Start the application
/opt/app/wallet-core