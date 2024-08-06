FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go get ./...

COPY main.go ./
COPY internal/ ./internal
COPY cmd/ ./cmd

RUN go build -o clread

ENTRYPOINT ["./clread"]
