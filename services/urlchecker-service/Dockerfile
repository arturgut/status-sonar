FROM golang:1.14.2

WORKDIR /go/app
COPY ./bin/urlchecker-linux-amd64 . 
COPY ./config.yaml .

ENV URL_CHECKER_SERVER_PORT=8091

EXPOSE 8091

CMD ["./urlchecker-linux-amd64"]
