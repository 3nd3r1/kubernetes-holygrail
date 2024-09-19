FROM golang:1.22-alpine AS build-stage

WORKDIR /usr/src/app

COPY . .
RUN adduser -S appuser && go mod download

WORKDIR /usr/src/app/cmd/app
RUN GCO_ENABLED=0 go build -o /usr/src/app/todo-broadcaster

FROM scratch

ENV PORT=8080

EXPOSE ${PORT}

WORKDIR /usr/src/app

COPY --from=build-stage /etc/passwd /etc/passwd
COPY --from=build-stage /usr/src/app/todo-broadcaster ./todo-broadcaster

USER appuser

ENTRYPOINT ["./todo-broadcaster"]
