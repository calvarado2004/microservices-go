FROM docker.io/golang:1.18.4-alpine as builder

RUN mkdir /app

COPY ./cmd /app/cmd

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/front-end

RUN CGO_ENABLED=0 go build -o frontEndApp ./cmd/web

RUN chmod +x /app/frontEndApp

FROM alpine:latest 

RUN mkdir /app

COPY ./cmd/web/templates /app/cmd/web/templates

COPY --from=builder /app/frontEndApp /app

CMD [ "/app/frontEndApp"]