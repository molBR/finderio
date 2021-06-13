FROM golang:1.14.3-alpine AS build
WORKDIR /go/src/finderio
COPY . .
RUN go get -d -v
RUN go install finderio
ENTRYPOINT /go/bin/finderio
EXPOSE 8080
