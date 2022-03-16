FROM golang:1.16 AS builder

Run mkdir /app
ADD . /app
WORKDIR /app