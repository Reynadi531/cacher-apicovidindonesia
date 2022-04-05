FROM golang:1.18-buster AS builder

RUN apt update && \
    apt install build-essential -y

WORKDIR /app

COPY . .

RUN make build

FROM debian:buster-slim

COPY --from=builder /app/bin/main /app/main
WORKDIR /app

RUN mkdir ./results

RUN chmod +x ./main

CMD ["./main"]
