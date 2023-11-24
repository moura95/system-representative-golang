FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/http_viewer/http_viewer.go

FROM alpine

COPY --from=builder /app/http_viewer /usr/local/bin/http_viewer

COPY --from=builder /app/.env /usr/local/bin/.env

COPY --from=builder /app/templates/ /usr/local/bin/templates/

EXPOSE 8000

WORKDIR /usr/local/bin/

ENTRYPOINT ["http_viewer"]
