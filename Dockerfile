FROM registry.cn-hongkong.aliyuncs.com/sync-dockerimage/go-build-utils:v3 AS cacher

RUN echo "build cache"

FROM cacher AS builder

COPY . /go/src/
WORKDIR /go/src/

RUN make get \
    && make error \
    && make ent \
    && make types \
    && make config \
    && make build

# build copy_middle
FROM debian:11-slim AS copy_middle

COPY --from=builder /go/src/bin/litmedia /srv/litmedia

FROM debian:11-slim AS http
WORKDIR /srv/

RUN apt-get update && apt-get install ca-certificates -y && rm -rf /tmp/* && rm -rf /var/lib/apt/lists/*

COPY --from=copy_middle /srv/litmedia /srv/app

CMD ["/srv/app"]