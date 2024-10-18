# builder image to build image
FROM golang:1.23.1-alpine3.20 as builder

ARG CI_COMMIT_BRANCH
ARG CI_COMMIT_SHA
ARG CI_PROJECT_URL
ARG BUILD_NUMBER

ENV CGO_ENABLED=0 \
    GOOS=linux
ADD . /build/
WORKDIR /build
RUN go build -o /build/locg-server -ldflags "-s -w -X locgame-mini-server/internal/version.RELEASE=${CI_COMMIT_BRANCH} -X locgame-mini-server/internal/version.COMMIT=${CI_COMMIT_SHA} -X locgame-mini-server/internal/version.REPO=${CI_PROJECT_URL} -X locgame-mini-server/internal/version.BUILD=${BUILD_NUMBER}" cmd/locgame-server/server.go
# RUN go build -o /build/log-test -ldflags "-s -w -X locgame-mini-server/internal/version.RELEASE=${CI_COMMIT_BRANCH} -X locgame-mini-server/internal/version.COMMIT=${CI_COMMIT_SHA} -X locgame-mini-server/internal/version.REPO=${CI_PROJECT_URL} -X locgame-mini-server/internal/version.BUILD=${BUILD_NUMBER}" cmd/metric-test/metric-test.go

# runtime image
FROM alpine:3.15
RUN addgroup -g 1000 locg && adduser -u 1000 -G locg -s /bin/sh -D locg && \
    apk --update upgrade && \
    apk add --update inotify-tools && \
    apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/*
COPY --chown=locg:locg --from=builder /build/locg-server /app/
# COPY --chown=locg:locg --from=builder /build/log-test /app/
COPY entrypoint.sh /docker-init.d/


WORKDIR /app
USER 1000
ENTRYPOINT [ "/docker-init.d/entrypoint.sh" ]
CMD [ "/app/locg-server" ]
