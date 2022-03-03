FROM golang:1.16-alpine AS build

ENV APP=./cmd/app
ENV BIN=/bin/echo
ENV PATH_ROJECT=${GOPATH}/src/git.repo.services.lenvendo.ru/grade-factor/echo
ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOFLAGS=-mod=vendor
ARG GITLAB_DEPLOYMENT_PRIVATE_KEY
ENV GITLAB_DEPLOYMENT_PRIVATE_KEY ${GITLAB_DEPLOYMENT_PRIVATE_KEY:-unknown}
ARG VERSION
ENV VERSION ${VERSION:-0.1.0}
ARG BUILD_TIME
ENV BUILD_TIME ${BUILD_TIME:-unknown}
ARG COMMIT
ENV COMMIT ${COMMIT:-unknown}


ADD https://github.com/golang-migrate/migrate/releases/download/v4.15.0/migrate.linux-amd64.tar.gz /tmp/
RUN cd  /tmp && ls && tar -xzf migrate.linux-amd64.tar.gz && chmod +x /tmp/migrate

WORKDIR ${PATH_ROJECT}
COPY . ${PATH_ROJECT}


#RUN apk add --no-cache git openssh-client ca-certificates gcc make alpine-sdk
#COPY --from=certs /certs/git.repo.services.lenvendo.ru.crt /usr/local/share/ca-certificates/git.repo.services.lenvendo.ru.ca.crt
#RUN update-ca-certificates

#RUN mkdir -p ~/.ssh \
#    && umask 0077 \
#    && echo "${GITLAB_DEPLOYMENT_PRIVATE_KEY}" > ~/.ssh/id_rsa \
#    && git config --global url."git@git.repo.services.lenvendo.ru:".insteadOf http://git.repo.services.lenvendo.ru/ \
#    && ssh-keyscan git.repo.services.lenvendo.ru >> ~/.ssh/known_hosts

WORKDIR ${PATH_ROJECT}
COPY . ${PATH_ROJECT}
CMD go test -v -race -timeout=5s ./...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w \
        -X git.repo.services.lenvendo.ru/grade-factor/tools/health.Version=${VERSION} \
        -X git.repo.services.lenvendo.ru/grade-factor/tools/health.Commit=${COMMIT} \
        -X git.repo.services.lenvendo.ru/grade-factor/tools/health.BuildTime=${BUILD_TIME}" \
    -a -installsuffix cgo -o ${BIN} ${APP}

COPY --from=build /tmp/migrate /bin/migrate
COPY --from=build /bin/echo /bin/echo

WORKDIR /migrations
COPY ./migrations /tmp/migrations

ENTRYPOINT ["/bin/echo"]
