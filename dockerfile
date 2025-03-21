FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache tzdata

RUN go build -o go-gin

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

COPY --from=builder /app/go-gin .

ENV TZ=Asia/Jakarta

EXPOSE 8000

ENTRYPOINT ["./go-gin"]
