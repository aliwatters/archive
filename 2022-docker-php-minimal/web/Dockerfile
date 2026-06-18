FROM nginx:alpine

EXPOSE 80

COPY ./web/nginx.conf /etc/nginx/nginx.conf
COPY ./web/http.js /etc/nginx/http.js
COPY ./web/conf /etc/nginx/conf.d

COPY ./src/html /app/html
