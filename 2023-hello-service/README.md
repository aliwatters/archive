# Hello Service

Intent - quick service for testing apitible installation of a typescript/node service.

See: https://hingehealth.atlassian.net/browse/NGMGMT-217

```
build out a simple node typescript micro-service
rest based, hapi/fastify/? - have a /hello end point
docker from a slim image, dockerfile
set up a pass through from basilisk/other service
heart beat end point - /status
has a watcher, lint etc – looking at this as being 'hh-rest-node-starter-kit' - later work

AC

deploy to aptible
ingest env vars for some config
returns "hello world" from the service - but only accessible internally
some other route (gateway service, basilisk passthrough?) passes the results through
end point should not be able to hit from outside our internal services
goes through build pipeline and is deployed
```