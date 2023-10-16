FROM ubuntu
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19

WORKDIR wkspace

ADD event-server /wkspace/event-server

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod -R 777 /wkspace/event-server

VOLUME ["/project"]

ENTRYPOINT ["/wkspace/event-server"]




# 运行容器

#docker build -t event-server:0.1.0 .
# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name event-server event-server:0.1.0



