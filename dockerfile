# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src
RUN swag init --parseDependency --parseInternal
RUN go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app
