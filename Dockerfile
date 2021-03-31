FROM golang:alpine

WORKDIR /app

COPY ./main.go ./

RUN go build -o main .

EXPOSE 3000

CMD ["/app/main"]