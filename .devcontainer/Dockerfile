# builder image to build image
FROM golang:1.22.4-alpine3.20 as builder


ENV CGO_ENABLED=0 \
    GOOS=linux\
    GO111MODULE=on 
ADD . /build/
WORKDIR /build
RUN pwd
RUN go build -o /build/locgame-jobs-service -ldflags "-s -w -X locgame-service/internal/version.RELEASE=development -X locgame-service/internal/version.COMMIT=local -X locgame-service/internal/version.REPO=local -X locgame-service/internal/version.BUILD=development" ./cmd/locgame-jobs-service/service.go
# RUN curl -o /usr/local/bin/git-credential-helper \
#     https://raw.githubusercontent.com/microsoft/vscode-dev-containers/master/script-library/git-credential-helper.sh \
#     && chmod +x /usr/local/bin/git-credential-helper
# RUN git config --global credential.helper "/usr/local/bin/git-credential-helper $@"


# runtime image
FROM alpine:3.15
RUN addgroup -g 1000 locg && adduser -u 1000 -G locg -s /bin/sh -D locg && \
    apk --update upgrade && \
    apk add --update inotify-tools && \
    apk add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/* && \
    pwd
COPY --chown=locg:locg --from=builder /build/locgame-jobs-service /app/
COPY ./cmd/locgame-jobs-service/build/entrypoint.sh /docker-init.d/



WORKDIR /app
USER 1000
ENTRYPOINT [ "/docker-init.d/entrypoint.sh" ]
CMD [ "/app/locgame-jobs-service" ]
