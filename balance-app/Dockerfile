FROM eclipse-temurin:23
MAINTAINER Josenaldo de Oliveira Matos Filho

RUN mkdir /opt/app
WORKDIR /opt/app
COPY build/libs/*.jar /opt/app/app.jar

EXPOSE 8080
CMD ["java", "-jar", "/opt/app/app.jar"]
