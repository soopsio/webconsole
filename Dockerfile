FROM shibingli/alpine:3.6
MAINTAINER Eric Shi <shibingli@realclouds.org>
ADD . /opt/app/
EXPOSE 8080
CMD ["/opt/app/run"]