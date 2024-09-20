kubectl delete deployment,service,ingress,pvc,secret,statefulset,cronjob,job -l project=todo-project --ignore-not-found=true
kubectl delete deployment,service,ingress,pvc,secret,statefulset,cronjob,job -l project=ping-pong --ignore-not-found=true
kubectl delete deployment,service,ingress,pvc,secret,statefulset,cronjob,job -l project=log-output --ignore-not-found=true
