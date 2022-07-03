FROM golang
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /docker-go-api
EXPOSE 8080
CMD [ "/docker-go-api" ]


