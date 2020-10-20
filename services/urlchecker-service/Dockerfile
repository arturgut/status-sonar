FROM golang:1.14.2

WORKDIR /go/app
COPY ../../tmp/urlchecker-linux-amd64 . 
COPY ./services/urlchecker-service/config.yaml .

ENV URL_CHECKER_SERVER_PORT=8091

EXPOSE 8091

CMD ["./urlchecker-linux-amd64"]
