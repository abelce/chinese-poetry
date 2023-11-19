FROM postgres:9.6.5-alpine
COPY ./{{lowerCase .Name}}db.sql /docker-entrypoint-initdb.d/{{lowerCase .Name}}db.sql