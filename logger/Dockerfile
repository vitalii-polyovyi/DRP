FROM golang:1.19-alpine


RUN apk add build-base

RUN mkdir /app

ADD . /app

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN go build -v -o main .

CMD ["/app/main"]
