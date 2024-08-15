FROM google/cloud-sdk:alpine

RUN apk add --update --no-cache postgresql-client

COPY ./backup-agent.sh /usr/src/app/backup-agent.sh
RUN chmod +x /usr/src/app/backup-agent.sh

ENTRYPOINT ["/usr/src/app/backup-agent.sh"]
