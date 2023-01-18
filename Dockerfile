FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o echo

FROM alpine

WORKDIR /usr/local/bin

COPY --from=builder /app/echo .

RUN chmod +x echo

ENV HOST "0.0.0.0:80"

EXPOSE 80

CMD ["echo"]
