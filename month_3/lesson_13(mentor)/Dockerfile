FROM golang:alpine

RUN mkdir app

ADD . /app
# COPY . /app
WORKDIR /app

RUN go build -o main ./cmd/main.go

CMD ["./main"]