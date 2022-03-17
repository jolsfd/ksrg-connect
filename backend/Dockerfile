FROM golang:1.18-alpine
RUN apk add build-base

WORKDIR /usr/src/app/backend

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]