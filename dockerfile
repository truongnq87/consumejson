FROM golang:1.14-alpine
RUN apk add --no-cache git
#ENV GO111MODULE=on
WORKDIR /go/src/app
#COPY . .

ENV GOFLAGS=-mod=vendor
ENV APP_USER app



#RUN go get github.com/gin-gonic/gin
#RUN go get github.com/dlclark/regexp2

ADD main.go /go/src/app
RUN go get -d -v ./...

#RUN go get ./...

# Build the Go app
RUN go build -o /go/src/app/go-consumejson-app .


# Set the Current Working Directory inside the container


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["/go/src/app/go-consumejson-app"]