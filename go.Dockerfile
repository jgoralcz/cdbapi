FROM golang:latest

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

LABEL owner = jgoralcz
LABEL serviceVersion = 1.0.0
LABEL description = "Golang Character Database but the container is big and has all of the code"

WORKDIR /build

# COPY go.mod .
# COPY go.sum .

COPY . .
RUN go mod download

# docs
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN go get -u github.com/swaggo/echo-swagger
RUN swag init

RUN go build -o main .

EXPOSE 8443

ENTRYPOINT ["./main"]
