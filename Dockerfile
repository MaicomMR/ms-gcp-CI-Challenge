FROM golang:1.12-alpine AS builder

WORKDIR /app
COPY main.go /app
COPY main_test.go /app

RUN go run main.go
RUN go test

FROM scratch

WORKDIR /app
COPY --from=builder /app .

CMD ["/main"]
