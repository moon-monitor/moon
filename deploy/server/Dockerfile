FROM ghcr.io/moon-monitor/moon:builder AS builder

ARG APP_NAME=""

COPY ../.. /moon

WORKDIR /moon

RUN make init
RUN make all
RUN make build app=${APP_NAME}

FROM ghcr.io/moon-monitor/moon:deploy

ARG APP_NAME=""

WORKDIR /moon

COPY --from=builder /moon/bin /moon/bin
COPY --from=builder /moon/cmd/${APP_NAME}/config /moon/config
COPY --from=builder /moon/i18n /moon/i18n
COPY --from=builder /moon/swagger /moon/swagger

RUN mv /moon/bin/${APP_NAME} /moon/bin/server

CMD ["./bin/server"]

