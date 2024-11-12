# Stage 1: Build the application
FROM gradle:8-jdk23 AS build
WORKDIR /app
COPY . .
RUN gradle build --no-daemon -x test

# Stage 2: Run the application
FROM eclipse-temurin:23-jre-alpine

WORKDIR /opt/app

RUN apk add --no-cache curl bash mysql-client 

COPY --from=build /app/build/libs/*.jar /opt/app/app.jar

COPY entrypoint.sh /opt/app/entrypoint.sh
RUN chmod +x /opt/app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/opt/app/entrypoint.sh"]
