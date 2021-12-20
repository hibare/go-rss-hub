# go-rss-hub


[![Go Report Card](https://goreportcard.com/badge/github.com/hibare/go-rss-hub)](https://goreportcard.com/report/github.com/hibare/go-container-status)
[![Docker Hub](https://img.shields.io/docker/pulls/hibare/go-rss-hub)](https://hub.docker.com/r/hibare/go-rss-hub)
[![Docker image size](https://img.shields.io/docker/image-size/hibare/go-rss-hub/latest)](https://hub.docker.com/r/hibare/go-rss-hub) 
[![GitHub issues](https://img.shields.io/github/issues/hibare/go-rss-hub)](https://github.com/hibare/go-rss-hub/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/hibare/go-rss-hub)](https://github.com/hibare/go-rss-hub/pulls)
[![GitHub](https://img.shields.io/github/license/hibare/go-rss-hub)](https://github.com/hibare/go-rss-hub/blob/main/LICENSE)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/hibare/go-rss-hub)](https://github.com/hibare/go-rss-hub/releases)


A REST API designed in go to check for a container status

## Getting Started

go-rss-hub is packaged as docker container. Docker image is available on [Docker Hub](https://hub.docker.com/r/hibare/go-rss-hub).

### Docker run

```shell
docker run -p 5000:5000 -e LISTEN_ADDR='0.0.0.0' hibare/go-rss-hub
```

### Docker Compose

```yml
version: "3"
services:
  go-rss-hub:
    image: hibare/go-rss-hub
    container_name: go-rss-hub
    hostname: go-rss-hub
    environment:
      - LISTEN_ADDR=0.0.0.0
    ports:
      - "5000:5000"
    logging:
      driver: "json-file"
      options:
        max-size: "500k"
        max-file: "5"

```
## Endpoints

1. Check health

```shell
/ping/
```

2. Docker hub rss
RSS feed for images published on Docker hub.

```shell
/docker/<user>/<repository>/tags/
```

Example:
```shell
> curl http://127.0.0.1:5000/docker/hibare/vt_ip_monitor/tags/


<?xml version="1.0" encoding="UTF-8"?><feed xmlns="http://www.w3.org/2005/Atom">
  <title>hibare/vt_ip_monitor | Docker Hub Images</title>
  <id>https://hub.docker.com/r/hibare/vt_ip_monitor</id>
  <updated>2021-12-20T10:49:50Z</updated>
  <subtitle>A Python 3 script to monitor your IP for malicious domains/URL.</subtitle>
  <link href="https://hub.docker.com/r/hibare/vt_ip_monitor"></link>
  <author>
    <name>hibare</name>
  </author>
  <entry>
    <title>hibare/vt_ip_monitor:latest</title>
    <updated>2020-01-06T18:41:43Z</updated>
    <id>tag:hub.docker.com,2020-01-06:/r/hibare/moni/tags?name=latest</id>
    <content type="html">Docker image ID: 81667473, Status: inactive</content>
    <link href="https://hub.docker.com/r/hibare/vt_ip_monitor/tags?name=latest" rel="alternate"></link>
    <summary type="html">A Python 3 script to monitor your IP for malicious domains/URL.</summary>
    <author>
      <name>hibare</name>
    </author>
  </entry>
</feed>
```

## Supported Environment Variables
| Variable | Description | Default Value |
| --------- | ----------- | ------------- |
| LISTEN_ADDR | IP address to bind to | 127.0.0.1 |
| LISTEN_PORT | Port to bind to | 5000 |
