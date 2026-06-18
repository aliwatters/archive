FROM nginx:alpine

EXPOSE 80
EXPOSE 443

RUN apk add ca-certificates 
COPY ./load-balancer/conf /etc/nginx/conf.d
COPY ./load-balancer/ssl /ssl
