#!/bin/bash
FROM golang:1.14-alpine

LABEL maintainer="Patiphan Chaiya <patiphann@gmail.com>"

WORKDIR /app

COPY ./main /app

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
CMD ["./main"]

EXPOSE 1323