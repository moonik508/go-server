FROM library/golang:1.20-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app
COPY . .

ENV GO111MODULE=on

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7
RUN swag init

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o /api-server .

FROM scratch
COPY --from=builder /api-server /api-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["/api-server"]
