# Run the app in a tiny Linux image
FROM alpine:latest

RUN mkdir /app
COPY mailerApp /app
COPY templates /templates

CMD [ "/app/mailerApp"]