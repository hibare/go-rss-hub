FROM golang:1.20.5-alpine AS base

# Build golang healthcheck binary
FROM base AS healthcheck

ARG VERSION=0.1.0

RUN wget -O - https://github.com/hibare/go-docker-healthcheck/archive/refs/tags/v${VERSION}.tar.gz |  tar zxf -

WORKDIR /go/go-docker-healthcheck-${VERSION}

RUN CGO_ENABLED=0 go build -o /bin/healthcheck

# Build main app
FROM base AS build

RUN apk --update add --no-cache ca-certificates openssl git tzdata \
    && update-ca-certificates

WORKDIR /src/

COPY . /src/

RUN CGO_ENABLED=0 go build -o /bin/go-rss-hub

# Generate final image
FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /bin/go-rss-hub /bin/go-rss-hub

COPY --from=healthcheck /bin/healthcheck /bin/healthcheck

HEALTHCHECK \
    --interval=30s \
    --timeout=3s \
    CMD ["healthcheck","http://localhost:5000/ping/"]

EXPOSE 5000

ENTRYPOINT ["/bin/go-rss-hub"]