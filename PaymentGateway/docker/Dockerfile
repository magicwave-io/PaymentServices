FROM pg_build as build
FROM torplusserviceregistry.azurecr.io/private/tor:latest
ENV BASE_NAME application
RUN apt-get update && \
    apt-get install -y curl gettext-base supervisor && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /opt/${BASE_NAME}/
ENV seed=SAGWKBJQ5J2AJAM5YTEUGSOEXUBOVNZJ6FNJDN7CBKZ543VRQ7XTDFIQ
ENV INITIALIZE_TOR=true
ENV INITIALIZE_PG=true
ENV INITIALIZE_IPFS=true
COPY --from=build  /opt/${BASE_NAME}/main_linux payment-gateway
COPY PaymentGateway/config.json config.json
COPY PaymentGateway/docker/config.json.tmpl config.json.tmpl
COPY PaymentGateway/docker/pg-docker-entrypoint.sh /pg-docker-entrypoint.sh
RUN chmod 755 /pg-docker-entrypoint.sh
COPY PaymentGateway/docker/supervisord.conf /etc/supervisor/conf.d/supervisord.conf
COPY PaymentGateway/docker/debug-start.sh /debug-start.sh 
ENTRYPOINT ["/debug-start.sh"]
