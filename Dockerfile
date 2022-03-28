FROM golang:1.18-alpine

RUN apk update && \
  apk add bash alpine-sdk --no-cache

WORKDIR /app

COPY . .

RUN make build

CMD [ "./bin/main" ]