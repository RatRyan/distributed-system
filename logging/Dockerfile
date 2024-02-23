FROM bitnami/prometheus:latest

WORKDIR /etc/prometheus

COPY prometheus.yaml .

EXPOSE  9090

CMD ["--config.file=/etc/prometheus/prometheus.yaml"]