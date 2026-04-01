FROM golang:1.25.0-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /out/orderapi ./cmd/orderapi

FROM alpine:3.21

RUN addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /out/orderapi /app/orderapi

EXPOSE 8081

USER app

ENTRYPOINT ["/app/orderapi"]
