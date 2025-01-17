# Wikipedia Reader

Simple deployment that at random intervals downloads a random Wikipedia article and servers it via ingress.

This project was meant for practicing initContainers and sidecars.

## To deploy

1. Define your cluster and set kubectl to use that context.

2. Run

```bash
make deploy
```

The deployment will be deployed to wikipedia-reader namespace.
You can access it at [wikipedia-reader.fbi.com:8081](http://wikipedia-reader.fbi.com:8081)

## To uninstall

Run

```bash
make uninstall
```
