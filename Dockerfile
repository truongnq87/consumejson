FROM golang:1.14-alpine
RUN apk add --no-cache git
WORKDIR /go/src/app
ENV GOFLAGS=-mod=vendor
ENV APP_USER app

ADD main.go /go/src/app
RUN go get -d -v ./...

# Build the Go app
RUN go build -o /go/src/app/go-consumejson-app .

# This container exposes port 8080 to the outside world
EXPOSE 8080

CMD ["/go/src/app/go-consumejson-app"]
