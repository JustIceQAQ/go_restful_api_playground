# build stage
FROM golang AS build-env
ADD . /src
RUN cd /src && go build -o app

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENV GIN_MODE="release"
ENTRYPOINT ./app
