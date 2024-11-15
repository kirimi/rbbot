FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/rbbot cmd/rb_bot/main.go

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=builder /app/rbbot /app/rbbot
CMD ["/app/rbbot"]