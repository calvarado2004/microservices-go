FROM docker.io/golang:1.19-alpine as builder

RUN mkdir /app

COPY ./cmd /app/cmd
COPY ./data /app/data
COPY ./logs /app/logs

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/log-service 

RUN go get github.com/go-chi/chi/v5 && go get github.com/go-chi/cors && go get go.mongodb.org/mongo-driver/mongo && go get go.mongodb.org/mongo-driver/mongo/options && go get google.golang.org/grpc

RUN CGO_ENABLED=0 go build -o loggerServiceApp ./cmd/api

RUN chmod +x /app/loggerServiceApp

FROM alpine:latest 

RUN mkdir /app

COPY --from=builder /app/loggerServiceApp /app

CMD [ "/app/loggerServiceApp"]