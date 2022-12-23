FROM openjdk:8u342

RUN adduser ctf && addgroup ctfgroup && adduser ctf ctfgroup
RUN mkdir -p /opt/app /home/ruoyi/logs && chown ctf:ctfgroup /home/ruoyi/logs
COPY wait-for /
COPY nc /usr/bin/
COPY ruoyi-admin.jar /opt/app

WORKDIR /opt/app

USER ctf
#CMD ["java", "-jar", "/opt/app/ruoyi-admin.jar"]

