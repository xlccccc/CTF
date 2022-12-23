FROM openjdk:8u342-slim

RUN adduser ctf && addgroup ctfgroup && adduser ctf ctfgroup
RUN mkdir /opt/app

COPY demo-0.0.1-SNAPSHOT.jar /opt/app
COPY flag /flag

WORKDIR /opt/app

USER ctf
CMD ["java", "-jar", "/opt/app/demo-0.0.1-SNAPSHOT.jar"]

