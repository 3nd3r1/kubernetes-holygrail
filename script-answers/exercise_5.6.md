# Exercise 5.6

```bash
3nd3r1@mypc$ kubectl apply -f hello1.yaml
3nd3r1@mypc$ kubectl get ksvc
3nd3r1@mypc$ curl -H "Host: hello.default.172.19.0.2.sslip.io" http://localhost:8081
Hello World!
3nd3r1@mypc$ kubectl apply -f hello2.yaml
3nd3r1@mypc$ kubectl get revisions
NAME          CONFIG NAME   GENERATION   READY   REASON   ACTUAL REPLICAS   DESIRED REPLICAS
hello-00001   hello         1            True             1                 1
hello-00002   hello         2            True             1                 1
3nd3r1@mypc$ curl -H "Host: hello.default.172.19.0.2.sslip.io" http://localhost:8081
Hello Pöpö!
3nd3r1@mypc$ curl -H "Host: hello.default.172.19.0.2.sslip.io" http://localhost:8081
Hello Pöpö!
3nd3r1@mypc$ curl -H "Host: hello.default.172.19.0.2.sslip.io" http://localhost:8081
Hello World!
```

Content of hello1.yaml:

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: hello
spec:
  template:
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          ports:
            - containerPort: 8080
          env:
            - name: TARGET
              value: "World"
```

Content of hello2.yaml:

```yaml
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: hello
spec:
  template:
    spec:
      containers:
        - image: ghcr.io/knative/helloworld-go:latest
          ports:
            - containerPort: 8080
          env:
            - name: TARGET
              value: "Pöpö"
  traffic:
  - latestRevision: true
    percent: 50
  - latestRevision: false
    percent: 50
    revisionName: hello-00001
```
