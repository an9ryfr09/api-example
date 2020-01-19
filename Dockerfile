FROM golang:1.13

LABEL author "an9ryfr09" \
    app "a6-api" \
    version="1.0"

WORKDIR /a6-api

ADD ./* /a6-api/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./main.go

EXPOSE 80

ENTRYPOINT ["./api"]