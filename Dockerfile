FROM golang:1.18-buster AS builder

RUN apt update && \
    apt install build-essential -y

WORKDIR /app

COPY . .

RUN make build

FROM alpine:latest

COPY --from=builder /app/bin/main /app/main
WORKDIR /app

CMD ["./main"]
