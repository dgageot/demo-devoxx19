FROM postgres:11.2-alpine

ENV POSTGRES_PASSWORD notsosecret
ENV POSTGRES_USER hexagons
ENV POSTGRES_DB hexagons

COPY init.sql /docker-entrypoint-initdb.d/
