FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /app/cmd/server/gophkeeper /app/cmd/server

