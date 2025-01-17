# Kubernetes Holygray

I use this repository for testing various Kubernetes features and concepts. Each project is a separate directory with its own README.md file.

## How to get a cluster?

All of the projects obviously require a Kubernetes cluster. Don't have one? No worries.
Just install [Docker](https://www.docker.com/get-started), [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), and [k3d](https://k3d.io/).

Then, run the following command:

```bash
bob create-cluster
```

This will create a small cluster that is actually a bunch of Docker containers.

Now you are all set! Enjoy the projects!

## Projects

-   [Ping And Log](ping-and-log/README.md)
-   [Todo Project](todo-project/README.md)
-   [Wikipedia Reader](wikipedia-reader/README.md)
-   [Dummy Site](dummy-site/README.md)
