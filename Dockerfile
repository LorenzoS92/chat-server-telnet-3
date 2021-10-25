# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /chat-server

ENV CHAT_PORT=":10000"
EXPOSE 10000

CMD [ "/chat-server" ]