# Minimal Docker Development Setup

Aim here is to provide a docker setup that matches closely with a production deployment. This will ensure that images work as intended.

```
load-balancer (ssl) --> web (rewrites, assets and proxy) --> php (8.1) --> database (to follow)
```

In production a managed load-balancer would be used.

Next steps:

1. add a mysql database (with utf-8)
2. rework into a kubernetes cluster

# Usage

Requires docker >20 -- check with `docker -v`, preferred install in ubuntu `sudo snap install docker`

```
cd load-balancer
./setup.sh

sudo vi /etc/hosts
# add dev.example.org

cd ..

docker-compose build
docker-compose up
```

Visit: https://dev.example.org/ and accept the "self-signed certificate"

To see an end with utf-8 from the mysql db see: https://dev.example.org/countries/us
