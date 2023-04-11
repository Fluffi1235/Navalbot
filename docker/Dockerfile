FROM postgres:latest

ENV POSTGRES_USER fluffi
ENV POSTGRES_PASSWORD nedji240696
ENV POSTGRES_DB navalbot

COPY init.sql /docker-entrypoint-initdb.d/

EXPOSE 32750