# Ping And Log

Ping and Log is a straightforward Kubernetes deployment with two components:

-   **Ping Pong**: An HTTP server that responds with "pong" to GET requests and tracks the total request count in a PostgreSQL database.
-   **Log Output**: A service that periodically retrieves the request count from Ping Pong, logs it to the console and displays it via HTTP server.

This project serves as a hands-on refresher on pod communication, volume management, and database integration within Kubernetes.

## How to deploy

To deploy you need [Linkerd](https://linkerd.io/2.17/getting-started/#step-1-install-the-cli)

1. Create a Kubernetes cluster and set the kubectl context to that cluster.

2. Run

```bash
make deploy
```

The project will be deployed to the namespace ping-and-log.

-   Log Output will be available at [ping-and-log.fbi.com:8081/](http://ping-and-log.fbi.com:8081).
-   Ping Pong will be available at [ping-and-log.fbi.com:8081/pingpong](http://ping-and-log.fbi.com:8081/pingpong).

## How to deploy using knative

1. Create a Kubernetes cluster with Knative and set the kubectl context to that cluster.

2. Run

```bash
make deploy-knative
```

## View service mesh

The project is using a service mesh called Linkerd

You can view that by running:

```bash
linkerd viz dashboard &
```

## How to uninstall

Run

```bash
make uninstall
```
