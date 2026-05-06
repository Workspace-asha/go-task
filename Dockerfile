FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go build -o api ./cmd/api

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/api /api
CMD ["/api"]