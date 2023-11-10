FROM golang:1.20-alpine

WORKDIR /app


COPY go.mod ./
RUN go mod download
COPY . ./


RUN go build -o /example-server cmd/main.go


CMD [ "/example-server" ]