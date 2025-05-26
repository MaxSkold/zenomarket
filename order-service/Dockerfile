FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod .
#COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/order-service ./main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates

RUN adduser -D -H -u 10001 app
USER 10001

COPY --from=builder /bin/order-service /app
EXPOSE 8083
ENTRYPOINT ["/app"]