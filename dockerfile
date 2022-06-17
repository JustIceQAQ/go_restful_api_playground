FROM golang:1.18.3-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app

EXPOSE 8080

ENTRYPOINT ["./app"]