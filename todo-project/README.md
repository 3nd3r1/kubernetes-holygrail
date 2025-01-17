# Todo Project

Todo project, as the name suggests, is a simple todo list application.
However, the deployment itself is far from simple. I have used almost all the Kubernetes features I know to deploy this application.

## How to deploy

Deploying the application in production mode requires my GCP encryption key. However, deploying the application in staging does not.

### Staging

Run

```bash
make deploy-staging
```

### Production

Run

```bash
make deploy-staging
```

### Development

Development means no Kubernetes cluster, but a Docker container

Run

```bash
docker compose up
```

## View service mesh

The project is using a service mesh called Linkerd

You can view that by running:

```bash
linkerd viz dashboard &
```

## How to uninstall

To uninstall the application, run

```bash
make uninstall
```
