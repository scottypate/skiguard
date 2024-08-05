FROM apache/superset:f11fa09

USER root

ARG SUPERSET_SECRET_KEY
ENV SUPERSET_SECRET_KEY=SUPERSET_SECRET_KEY

RUN apt update && apt install -y jq cron vim systemd

RUN cd /usr/local && \
    curl -O https://dl.google.com/go/go1.22.5.linux-arm64.tar.gz && \
    tar -xvf go1.22.5.linux-arm64.tar.gz && \
    rm go1.22.5.linux-arm64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin && \
    go version

RUN pip install --upgrade pip && \
    pip install duckdb-engine

COPY ./db /home/skiguard/db

RUN superset fab create-admin --username admin --password admin --firstname admin --lastname admin --email admin@admin.com && \
    superset db upgrade && \
    superset init && \
    superset import-datasources -p /home/skiguard/db/superset/datasources.zip && \
    superset import-dashboards -p /home/skiguard/db/superset/dashboard.zip -u admin

RUN echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

WORKDIR /home/skiguard

COPY .docker/entrypoint.sh /home/skiguard/entrypoint.sh
COPY ./docs/images/skiguard.png /app/superset/static/assets/images/skiguard.png
COPY ./bin /home/skiguard/bin
COPY .docker/superset_config.py /app/superset/config.py
COPY .docker/skiguard-cron /etc/cron.d/skiguard-cron
COPY ./internal/api/alert/templates/ /home/skiguard/internal/api/alert/templates/

RUN crontab /etc/cron.d/skiguard-cron && touch /var/log/cron.log

ENTRYPOINT ["/home/skiguard/entrypoint.sh"]
