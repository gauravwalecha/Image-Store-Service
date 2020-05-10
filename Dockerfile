FROM golang:latest
ADD . /go/
WORKDIR /go
RUN go get -d ./...
EXPOSE 8080
CMD go run main.go