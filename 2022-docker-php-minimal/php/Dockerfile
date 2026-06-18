FROM php:8.1-fpm-alpine

ENV PHP_OPCACHE_VALIDATE_TIMESTAMPS="0" \
    PHP_OPCACHE_MAX_ACCELERATED_FILES="10000" \
    PHP_OPCACHE_MEMORY_CONSUMPTION="192" \
    PHP_OPCACHE_MAX_WASTED_PERCENTAGE="10"

RUN docker-php-ext-install opcache

RUN docker-php-ext-install mysqli

COPY ./php/opcache.ini /usr/local/etc/php/conf.d/opcache.ini

COPY ./src /app
