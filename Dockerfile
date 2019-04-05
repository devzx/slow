FROM golang:1.12.1-alpine3.9 as builder
WORKDIR /app
COPY . .
RUN go build *.go

FROM alpine:3.9.2
COPY --from=builder /app/slow /usr/bin

CMD ["./slow"]
