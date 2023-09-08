FROM ubuntu:latest

VOLUME /root/.hostker
## 安装依赖
RUN apt-get update && \
    apt-get install -y curl

RUN curl -o /tmp/install.sh -L https://raw.githubusercontent.com/csvwolf/hostker-ddns/master/install.sh && \
    chmod +x /tmp/install.sh && \
    /tmp/install.sh
ENTRYPOINT ["hostker-ddns", "start"]