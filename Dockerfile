FROM alpine:3.19

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

WORKDIR /app
COPY build/glance-$TARGETOS-$TARGETARCH${TARGETVARIANT} /app/glance
COPY glance.yml /app/glance.yml

# Create assets directory
RUN mkdir -p /app/assets
# COPY path/to/your/assets /app/assets

EXPOSE 8080
ENV PORT 8080
ENTRYPOINT ["/app/glance", "--config", "/app/glance.yml"]
