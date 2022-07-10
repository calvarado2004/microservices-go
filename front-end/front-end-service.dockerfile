FROM golang:1.18-alphine as builder

RUN mkdir /app
COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o frontEndApp ./cmd/api

RUN chmod +x /app/frontEndApp

FROM alpine:latest 

RUN mkdir /app

COPY --from=builder /app/frontEndApp /app

CMD [ "/app/frontEndApp"]