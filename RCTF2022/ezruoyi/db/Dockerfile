FROM mysql:8.0.26

ENV MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=ry
COPY my.cnf /etc/mysql/
COPY flag.sql /docker-entrypoint-initdb.d/
COPY ry_20210924.sql /docker-entrypoint-initdb.d/
COPY quartz.sql /docker-entrypoint-initdb.d/