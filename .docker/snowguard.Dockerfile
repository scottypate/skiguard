FROM apache/superset:latest

USER root

ARG SUPERSET_SECRET_KEY
ENV SUPERSET_SECRET_KEY=SUPERSET_SECRET_KEY

RUN apt update && apt install -y jq cron vim

RUN cd /usr/local && \
    curl -O https://dl.google.com/go/go1.22.5.linux-arm64.tar.gz && \
    tar -xvf go1.22.5.linux-arm64.tar.gz && \
    rm go1.22.5.linux-arm64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin && \
    go version

RUN pip install --upgrade pip && \
    pip install duckdb-engine

RUN superset fab create-admin --username admin --password admin --firstname admin --lastname admin --email admin@admin.com && \
    superset db upgrade && \
    superset init

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

WORKDIR /home/snowguard

COPY .docker/entrypoint.sh /home/snowguard/entrypoint.sh
COPY ./bin /home/snowguard/bin
COPY ./db /home/snowguard/db
COPY .docker/superset_config.py /app/superset/config.py
COPY .docker/snowguard-cron /etc/cron.d/snowguard-cron

RUN crontab /etc/cron.d/snowguard-cron && touch /var/log/cron.log

ENTRYPOINT ["/home/snowguard/entrypoint.sh"]
