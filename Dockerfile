FROM golang:1.15-alpine

WORKDIR /app

COPY . .
RUN go mod download

WORKDIR /app/server/cmd
RUN go build -o server .

EXPOSE 1337

CMD ["/app/server/cmd/server"]
