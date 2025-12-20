FROM golang:alpine3.23

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum . /

RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 8081

CMD ["air"]

