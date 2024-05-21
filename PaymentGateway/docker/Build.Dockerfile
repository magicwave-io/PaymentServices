FROM golang:1.16.8 as build
ENV BASE_NAME application
WORKDIR /opt/${BASE_NAME}/go-payment-service
ARG PG_VERSION=devel
ENV PG_VERSION $PG_VERSION
ARG COMMIT_HASH=hash
ENV COMMIT_HASH $COMMIT_HASH
ARG BUILD_DATE=date
ENV BUILD_DATE $BUILD_DATE
RUN apt-get update && apt-get install -y build-essential manpages-dev make
COPY ./payment-gateway-webui/go.mod ./payment-gateway-webui/go.mod
COPY ./payment-gateway-webui/go.sum ./payment-gateway-webui/go.sum
COPY ./PaymentGateway/go.mod ./PaymentGateway/go.mod
COPY ./PaymentGateway/go.sum ./PaymentGateway/go.sum
COPY ./ProtocolCustomizations/go.mod ./ProtocolCustomizations/go.mod
RUN cd PaymentGateway && go mod download

COPY . .
WORKDIR /opt/${BASE_NAME}/go-payment-service/PaymentGateway
RUN make build_linux_docker
