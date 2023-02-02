FROM node:8.1.2
COPY files/start.sh /start.sh
RUN chmod +x  /start.sh
RUN mkdir /app && \
    chown -R root:root /app && \
    chmod -R 755 /app
COPY src /app/
RUN cd /app && \
    chmod +x /app/bin/www


WORKDIR /app

CMD /start.sh