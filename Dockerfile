FROM centos:latest

LABEL alphay "alphay@y-alpha.com"

WORKDIR /root

ADD ./main ./main
ADD ./static/ ./static/
ADD ./assets/ ./assets/
ADD ./admin.json ./admin.json
ADD ./config.json ./config.json
ADD ./count.gob ./count.gob
ADD ./plugin.so ./plugin.so
RUN chmod 755 main

EXPOSE 8001

ENTRYPOINT  ["./main"]
