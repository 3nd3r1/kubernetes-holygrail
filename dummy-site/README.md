# Dummy Site

The Dummy Site project introduces a Custom Resource Definition (CRD) designed to take any website URL and serve it through an ingress.
The project comprises two key components: the CRD itself, which defines the desired site configuration, and a controller that manages the lifecycle of the site, ensuring it's served seamlessly via ingress.

## How to deploy

1. Create a Kubernetes cluster and set the kubectl context to that cluster.

2. Run

```bash
make deploy
```

This will deploy the controller and create the CRD

3. To create an example object run

```bash
make deploy-example
```

To access it, you have to port-forward the pod.

## How to uninstall

Run

```bash
make uninstall
```
