# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app/matchingAppMatchingService

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /docker-app-matching-service

EXPOSE 8080

CMD [ "/docker-app-matching-service" ]