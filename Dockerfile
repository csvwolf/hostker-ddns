FROM ubuntu:latest

VOLUME $HOME/.hostker/ddns_config.yaml

# 安装依赖
RUN apt-get update && \
    apt-get install -y curl

RUN curl -o /tmp/install.sh -L https://github.com/csvwolf/hostker-ddns/raw/master/build.sh && \
    chmod +x /tmp/install.sh && \
    /tmp/install.sh
CMD ["hostker-ddns", "start"]
