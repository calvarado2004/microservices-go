FROM docker.io/golang:1.18.4-alpine as builder

RUN mkdir /app

COPY ./cmd /app/cmd

COPY ./event /app/event
COPY ./logs /app/logs

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/broker-service && go get github.com/go-chi/chi/v5 && go get github.com/go-chi/cors && go get github.com/rabbitmq/amqp091-go && go get google.golang.org/grpc

RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

RUN chmod +x /app/brokerApp

FROM alpine:latest 

RUN mkdir /app

COPY --from=builder /app/brokerApp /app

CMD [ "/app/brokerApp"]