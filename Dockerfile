FROM reg.miz.so/library/centos:7
MAINTAINER "Maizuo glaucus <bob@hyx.com>"

RUN rm /etc/localtime && ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
ADD ./remoteserver /usr/local/glaucus

WORKDIR /usr/local/glaucus
