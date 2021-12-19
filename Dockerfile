FROM golang:1.17.5-alpine AS base

# Build golang healthcheck binary
FROM base AS healthcheck

ARG VERSION=0.1.0

RUN wget -O - https://github.com/hibare/go-docker-healthcheck/archive/refs/tags/v${VERSION}.tar.gz |  tar zxf -

WORKDIR /go/go-docker-healthcheck-${VERSION}

RUN CGO_ENABLED=0 go build -o /bin/healthcheck

# Build main app
FROM base AS build

WORKDIR /src/

COPY . /src/

RUN CGO_ENABLED=0 go build -o /bin/docker_hub_rss

# Generate final image
FROM scratch

COPY --from=build /bin/docker_hub_rss /bin/docker_hub_rss

COPY --from=healthcheck /bin/healthcheck /bin/healthcheck

HEALTHCHECK \
    --interval=30s \
    --timeout=3s \
    CMD ["healthcheck","http://localhost:5000/ping/"]

EXPOSE 5000

ENTRYPOINT ["/bin/docker_hub_rss"]