kubectl delete namespace applications --ignore-not-found=true

kubectl delete deployment,service,ingress,pvc,secret,statefulset,cronjob,job -l project=todo-project --ignore-not-found=true
