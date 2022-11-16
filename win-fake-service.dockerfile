# FROM golang:1.19-alpine as builder

# RUN mkdir /app

# COPY . /app

# WORKDIR /app

# RUN CGO_ENABLED=0 go build -o fk_service main.go

# RUN chmod +x /app/fk_service

FROM alpine:latest

RUN mkdir /app

COPY /dist/fk_service /app

CMD [ "/app/fk_service" ]