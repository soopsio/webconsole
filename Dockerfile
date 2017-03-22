FROM alpine:latest

MAINTAINER Eric Shi <postmaster@apibox.club>

RUN apk --update --no-cache add curl

RUN mkdir -p /data/apibox
RUN curl -L 'https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz' | tar -zx -C /usr/local
ENV PATH /usr/local/go/bin:$PATH
ENV GOPATH /data/apibox/
ADD . /data/apibox
RUN cd /data/apibox/src/apibox.club/apibox/ && go install
EXPOSE 8080

CMD ["/data/apibox/bin/apibox","start"]