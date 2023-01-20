FROM ubuntu:18.04

COPY java /dubbo/java/


RUN chmod +x /dubbo/java/jdk1.8.0_202/bin/java



RUN sed -i "s/http:\/\/archive.ubuntu.com/http:\/\/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list && sed -i "s/http:\/\/security.ubuntu.com/http:\/\/mirrors.tuna.tsinghua.edu.cn/g" /etc/apt/sources.list
RUN apt-get update
RUN apt-get install python3 curl -y
COPY start.sh /start.sh
RUN chmod +x /start.sh
RUN echo "flag{test}">/flag
RUN chmod 755 /flag
WORKDIR /
EXPOSE 20880

ENTRYPOINT ["/start.sh"]