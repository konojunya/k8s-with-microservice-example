FROM golang:1.24.0

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go build -o server .

CMD ["./server"]

EXPOSE 8082