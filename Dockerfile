FROM golang:1.8.0-alpine

MAINTAINER Eric Shi <postmaster@apibox.club>

ADD . /go/
RUN  go install -v

EXPOSE 8080

CMD ["/go/bin/apibox","start"]