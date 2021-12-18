ARG BASE_BUILDER_IMAGE
ARG BASE_GO_IMAGE

FROM ${BASE_BUILDER_IMAGE} AS lib_builder
WORKDIR /foundry
RUN apt-get update -y && apt-get install -y build-essential cmake git
RUN git clone https://github.com/jgarff/rpi_ws281x.git \
  && cd rpi_ws281x \ 
  && mkdir build \
  && cd build \ 
  && cmake -D BUILD_SHARED=OFF -D BUILD_TEST=OFF .. \
  && cmake --build . \
  && make install

FROM ${BASE_GO_IMAGE} AS base_go
COPY --from=lib_builder /usr/local/lib/libws2811.a /usr/local/lib/
COPY --from=lib_builder /usr/local/include/ws2811 /usr/local/include/ws2811
RUN go get github.com/rpi-ws281x/rpi-ws281x-go

RUN apk update && apk upgrade
RUN apk add --no-cache bash git openssh autoconf automake libtool gettext gettext-dev make g++ texinfo curl

ARG DEVELOPER_UID
RUN adduser -s /bin/sh -u ${DEVELOPER_UID} -D developer
USER developer

WORKDIR /go/src/${APP_NAME}