FROM mysql:8.0.26

ENV MYSQL_ROOT_PASSWORD=root MYSQL_DATABASE=ezbypass
COPY users.sql /docker-entrypoint-initdb.d/users.sql
