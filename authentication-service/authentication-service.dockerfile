FROM docker.io/golang:1.19-alpine as builder

RUN mkdir /app
COPY ./cmd /app/cmd
COPY ./data /app/data

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/authentication 
RUN go get github.com/jackc/pgconn && go get github.com/jackc/pgx/v4 && go get github.com/go-chi/cors && go get github.com/go-chi/chi/v5
RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x /app/authApp

FROM alpine:latest 

RUN mkdir /app

COPY --from=builder /app/authApp /app

CMD [ "/app/authApp"]
