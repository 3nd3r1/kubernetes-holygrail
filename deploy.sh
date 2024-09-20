export SOPS_AGE_KEY_FILE=./key.txt

# Deploy argo-rollouts
kubectl create namespace argo-rollouts --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml

# Deploy argocd
kubectl create namespace argocd --dry-run=client -o yaml | kubectl apply -f -
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# Deploy ping-pong and log-output
kubectl create namespace applications --dry-run=client -o yaml | kubectl apply -f -
kubectl config set-context --current --namespace=applications
sops --decrypt ./ping-pong/manifests/secrets/secret.enc.yaml | kubectl apply -f -
kubectl apply -k ./ping-pong/
kubectl apply -k ./log-output/

# Deploy todo-project
#kubectl config set-context --current --namespace=default
#helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats
#sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -f -
#kubectl apply -k ./todo-project/
