FROM golang:alpine

WORKDIR /app

COPY ./ ./

RUN go build -o main .

EXPOSE 3000

CMD ["/app/main"]