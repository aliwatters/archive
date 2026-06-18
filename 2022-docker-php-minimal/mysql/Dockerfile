FROM mysql:5.7

COPY ./mysql/data/countries.sql /docker-entrypoint-initdb.d/init.sql

EXPOSE 3306