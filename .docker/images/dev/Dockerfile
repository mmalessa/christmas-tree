ARG BASE_GO_IMAGE
FROM ${BASE_GO_IMAGE} AS base

ARG DEVELOPER_UID
RUN adduser -s /bin/sh -u ${DEVELOPER_UID} -D developer
USER developer

RUN apk update && apk upgrade
RUN apk add --no-cache bash git openssh autoconf automake libtool gettext gettext-dev make g++ texinfo curl

WORKDIR /go/src/${APP_NAME}
