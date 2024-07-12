FROM golang:1.22-alpine AS build-stage

WORKDIR /usr/src/app

COPY . .
RUN adduser -S appuser && go mod download

WORKDIR /usr/src/app/cmd/writer
RUN GCO_ENABLED=0 go build -o /usr/src/app/log-output-writer

FROM scratch

WORKDIR /usr/src/app

COPY --from=build-stage /etc/passwd /etc/passwd
COPY --from=build-stage /usr/src/app/log-output-writer ./log-output-writer

USER appuser

ENTRYPOINT ["./log-output-writer"]
