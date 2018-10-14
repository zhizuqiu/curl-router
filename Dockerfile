FROM docker.io/alpine:latest

COPY curl-router /

ENTRYPOINT ["/curl-router"]