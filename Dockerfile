FROM golang:1.8.0-alpine

MAINTAINER Eric Shi <shibingli@realclouds.org>
ADD . /go/
RUN cd /go/src/apibox.club/apibox/ && go install -v
EXPOSE 8080

CMD ["/go/bin/apibox","start"]