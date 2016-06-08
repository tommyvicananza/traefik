
<p align="center">
<img src="docs/img/traefik.logo.png" alt="Træfɪk" title="Træfɪk" />
</p>
## Important first aff all checkout `github.com/tommyvicananza/go-marathon` to traefik branch
[![Build Status](https://travis-ci.org/containous/traefik.svg?branch=master)](https://travis-ci.org/containous/traefik)
[![Docs](https://img.shields.io/badge/docs-current-brightgreen.svg)](https://docs.traefik.io)
[![Go Report Card](https://goreportcard.com/badge/kubernetes/helm)](http://goreportcard.com/report/containous/traefik)
[![Image Layer](https://badge.imagelayers.io/traefik:latest.svg)](https://imagelayers.io/?images=traefik)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/containous/traefik/blob/master/LICENSE.md)
[![Join the chat at https://traefik.herokuapp.com](https://img.shields.io/badge/style-register-green.svg?style=social&label=Slack)](https://traefik.herokuapp.com)
[![Twitter](https://img.shields.io/twitter/follow/traefikproxy.svg?style=social)](https://twitter.com/intent/follow?screen_name=traefikproxy)


Træfɪk is a modern HTTP reverse proxy and load balancer made to deploy microservices with ease.
It supports several backends ([Docker](https://www.docker.com/), [Swarm](https://docs.docker.com/swarm), [Mesos/Marathon](https://mesosphere.github.io/marathon/), [Kubernetes](http://kubernetes.io/), [Consul](https://www.consul.io/), [Etcd](https://coreos.com/etcd/), [Zookeeper](https://zookeeper.apache.org), [BoltDB](https://github.com/boltdb/bolt), Rest API, file...) to manage its configuration automatically and dynamically.

## Overview

Imagine that you have deployed a bunch of microservices on your infrastructure. You probably used a service registry (like etcd or consul) and/or an orchestrator (swarm, Mesos/Marathon) to manage all these services.
If you want your users to access some of your microservices from the Internet, you will have to use a reverse proxy and configure it using virtual hosts or prefix paths:

- domain `api.domain.com` will point the microservice `api` in your private network
- path `domain.com/web` will point the microservice `web` in your private network
- domain `backoffice.domain.com` will point the microservices `backoffice` in your private network, load-balancing between your multiple instances

But a microservices architecture is dynamic... Services are added, removed, killed or upgraded often, eventually several times a day.

Traditional reverse-proxies are not natively dynamic. You can't change their configuration and hot-reload easily.

Here enters Træfɪk.

![Architecture](docs/img/architecture.png)

Træfɪk can listen to your service registry/orchestrator API, and knows each time a microservice is added, removed, killed or upgraded, and can generate its configuration automatically.
Routes to your services will be created instantly.

Run it and forget it!




## Features

- [It's fast](http://docs.traefik.io/benchmarks)
- No dependency hell, single binary made with go
- Rest API
- Multiple backends supported: Docker, Mesos/Marathon, Consul, Etcd, and more to come
- Watchers for backends, can listen change in backends to apply a new configuration automatically
- Hot-reloading of configuration. No need to restart the process
- Graceful shutdown http connections
- Circuit breakers on backends
- Round Robin, rebalancer load-balancers
- Rest Metrics
- [Tiny](https://imagelayers.io/?images=traefik) [official](https://hub.docker.com/r/_/traefik/) docker image included
- SSL backends support
- SSL frontend support (with SNI)
- Clean AngularJS Web UI
- Websocket support
- HTTP/2 support
- Retry request if network error
- [Let's Encrypt](https://letsencrypt.org) support (Automatic HTTPS)

## Demo


Here is a talk (in french) given by [Emile Vauge](https://github.com/emilevauge) at the [Devoxx France 2016](http://www.devoxx.fr) conference.
You will learn fundamental Træfɪk features and see some demos with Docker, Mesos/Marathon and Lets'Encrypt.

[![Traefik Devoxx France](http://img.youtube.com/vi/QvAz9mVx5TI/0.jpg)](http://www.youtube.com/watch?v=QvAz9mVx5TI)

## Web UI

You can access to a simple HTML frontend of Træfik.

![Web UI Providers](docs/img/web.frontend.png)
![Web UI Health](docs/img/traefik-health.png)

## Plumbing

- [Oxy](https://github.com/vulcand/oxy): an awesome proxy library made by Mailgun guys
- [Gorilla mux](https://github.com/gorilla/mux): famous request router
- [Negroni](https://github.com/codegangsta/negroni): web middlewares made simple
- [Manners](https://github.com/mailgun/manners): graceful shutdown of http.Handler servers
- [Lego](https://github.com/xenolf/lego): the best [Let's Encrypt](https://letsencrypt.org) library in go

## Quick start

- The simple way: grab the latest binary from the [releases](https://github.com/containous/traefik/releases) page and just run it with the [sample configuration file](https://raw.githubusercontent.com/containous/traefik/master/traefik.sample.toml):

```shell
./traefik --configFile=traefik.toml
```

- Use the tiny Docker image:

```shell
docker run -d -p 8080:8080 -p 80:80 -v $PWD/traefik.toml:/etc/traefik/traefik.toml traefik
```

- From sources:

```shell
git clone https://github.com/containous/traefik
rm -Rf static/
make generate-webui
go generate
go build
```

## Documentation

You can find the complete documentation [here](https://docs.traefik.io).

## Contributing

Please refer to [this section](.github/CONTRIBUTING.md).

## Support

You can join [![Join the chat at https://traefik.herokuapp.com](https://img.shields.io/badge/style-register-green.svg?style=social&label=Slack)](https://traefik.herokuapp.com) to get basic support.
If you prefer a commercial support, please contact [containo.us](https://containo.us) by mail: <mailto:support@containo.us>.

## Træfɪk here and there

These projects use Træfɪk internally. If your company uses Træfɪk, we would be glad to get your feedback :) Contact us on [![Join the chat at https://traefik.herokuapp.com](https://img.shields.io/badge/style-register-green.svg?style=social&label=Slack)](https://traefik.herokuapp.com)

- Project [Mantl](https://mantl.io/) from Cisco

![Web UI Providers](docs/img/mantl-logo.png)
> Mantl is a modern platform for rapidly deploying globally distributed services. A container orchestrator, docker, a network stack, something to pool your logs, something to monitor health, a sprinkle of service discovery and some automation.

- Project [Apollo](http://capgemini.github.io/devops/apollo/) from Cap Gemini

![Web UI Providers](docs/img/apollo-logo.png)
> Apollo is an open source project to aid with building and deploying IAAS and PAAS services. It is particularly geared towards managing containerized applications across multiple hosts, and big data type workloads. Apollo leverages other open source components to provide basic mechanisms for deployment, maintenance, and scaling of infrastructure and applications.

## Partners

[![Zenika](docs/img/zenika.logo.png)](https://zenika.com)

Zenika is one of the leading providers of professional Open Source services and agile methodologies in
Europe. We provide consulting, development, training and support for the world’s leading Open Source
software products.


[![Asteris](docs/img/asteris.logo.png)](https://aster.is)

Founded in 2014, Asteris creates next-generation infrastructure software for the modern datacenter. Asteris writes software that makes it easy for companies to implement continuous delivery and realtime data pipelines. We support the HashiCorp stack, along with Kubernetes, Apache Mesos, Spark and Kafka. We're core committers on mantl.io, consul-cli and mesos-consul.

## Maintainers
- Emile Vauge [@emilevauge](https://github.com/emilevauge)
- Vincent Demeester [@vdemeester](https://github.com/vdemeester)
- Samuel Berthe [@samber](https://github.com/samber)

## Credits

Kudos to [Peka](http://peka.byethost11.com/photoblog/) for his awesome work on the logo ![logo](docs/img/traefik.icon.png)
