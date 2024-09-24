FROM golang:latest

RUN apt-get update && apt-get install -y tzdata
ENV TZ="Europe/Vienna"

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build .

ENTRYPOINT [ "go", "run", "." ]