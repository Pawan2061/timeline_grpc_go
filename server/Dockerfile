FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080 50051

CMD ["./main"]
