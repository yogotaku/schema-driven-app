FROM golang:1.20-alpine

WORKDIR /home
RUN go install github.com/cosmtrek/air@latest

COPY ./go.mod /home
COPY ./go.sum /home
RUN go mod download
