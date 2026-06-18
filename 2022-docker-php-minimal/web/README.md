# Web

Docker based nginx server for handling rewrites and proxies.

Uses: https://www.cnginx.com/docs/njs/reference.html - to do lowercase redirection of requests.

Custom `nginx.conf` here is minor update (enables the njs module), which is included in the `nginx:alpine` image.
