FROM golang:latest

RUN apt-get update && \
    apt-get -qy full-upgrade && \
    apt-get install -qy curl && \
    apt-get install -qy curl && \
    curl -sSL https://get.docker.com/ | sh

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build .

ENTRYPOINT [ "go", "run", "." ]