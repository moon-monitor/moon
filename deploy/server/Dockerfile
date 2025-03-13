FROM ghcr.io/moon-monitor/moon:builder AS builder

ARG APP_NAME=""

COPY ../.. /moon

WORKDIR /moon

RUN make init
RUN make all
RUN make build app=${APP_NAME}

FROM debian:stable-slim

ARG APP_NAME=""

COPY --from=builder /moon/bin /moon/bin
COPY --from=builder /moon/config /moon/config
COPY --from=builder /moon/i18n /moon/i18n
COPY --from=builder /moon/swagger /moon/swagger

RUN mv /moon/bin/${APP_NAME} /moon/bin/server

WORKDIR /moon

RUN apt update && apt install -y ca-certificates curl wget vim && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

CMD ["./bin/server"]

