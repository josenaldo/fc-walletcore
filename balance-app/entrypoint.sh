#!/bin/bash

# Wait until MySQL is ready
echo "Waiting for MySQL to be ready..."

until mysqladmin ping -h "mysql-balance" --silent; do
  echo "Waiting for mysql-balance..."
  sleep 2
done

echo "MySQL is up and running."

# Start the application
java -jar -Dspring.profiles.active=prod /opt/app/app.jar