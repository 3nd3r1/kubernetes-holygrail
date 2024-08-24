```bash
3nd3r1@mypc$ helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
3nd3r1@mypc$ helm repo add stable https://charts.helm.sh/stable
3nd3r1@mypc$ kubectl create namespace prometheus
3nd3r1@mypc$ helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus
3nd3r1@mypc$ kubectl -n prometheus port-forward prometheus-kube-prometheus-stack-1724-prometheus-0 9090:9090
```

The query
```
count(kube_pod_info{created_by_kind="StatefulSet",namespace="prometheus"})
```

I get 3 as the result?
