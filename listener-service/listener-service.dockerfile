FROM docker.io/golang:1.18.4-alpine as builder

RUN mkdir /app

COPY main.go /app/main.go
COPY ./event /app/event

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/listener-service 

RUN go get github.com/rabbitmq/amqp091-go

RUN CGO_ENABLED=0 go build -o listenerApp ./event/

RUN chmod +x /app/listenerApp


FROM alpine:latest 

RUN mkdir /app

COPY --from=builder /app/listenerApp /app

CMD [ "/app/listenerApp"]