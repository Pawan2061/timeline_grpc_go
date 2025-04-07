FROM golang:1.23

WORKDIR /app

COPY . .

RUN cd server && GOOS=linux GOARCH=amd64 go build -o main .

EXPOSE 8080 50051

WORKDIR /app/server

CMD ["./main"]
