FROM docker.io/golang:1.18.4-alpine as builder

RUN mkdir /app

COPY ./cmd /app/cmd

COPY ./templates /app/templates

WORKDIR /app

RUN go mod init calvarado2004/microservices-go/mailer-service 

RUN go get github.com/go-chi/chi/v5 && go get github.com/go-chi/cors && go get github.com/xhit/go-simple-mail/v2 && go get github.com/vanng822/go-premailer/premailer

RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api

RUN chmod +x /app/mailerApp

FROM alpine:latest 

WORKDIR /app

COPY ./templates /app/templates

COPY --from=builder /app/mailerApp /app

CMD [ "/app/mailerApp"]