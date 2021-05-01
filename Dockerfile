FROM ubuntu:latest
WORKDIR /app

EXPOSE 8080

COPY echo-webserver .

CMD [ "./echo-webserver" ]
