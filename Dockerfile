FROM golang:1.8.3-alpine

MAINTAINER Kalaiarasan skalaiarasan27@gmail.com

COPY .  /usr/local/go/src/go-tutorial/

RUN apk add --update git

RUN mkdir -p /usr/local/go/src/vendor/github.com/gorilla && cd /usr/local/go/src/vendor/github.com/gorilla && git clone https://github.com/gorilla/mux.git

RUN go install go-tutorial

CMD ["go-tutorial"]
