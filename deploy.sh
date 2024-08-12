export SOPS_AGE_KEY_FILE=./key.txt

#kubectl create namespace applications
#sops --decrypt ./ping-pong/manifests/secrets/secret.enc.yaml | kubectl apply -f -
#kubectl apply -f ./ping-pong/manifests/
#kubectl apply -f ./log-output/manifests/

kubectl apply -f ./todo-project/manifests/namespace.yaml
sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -f -
kubectl apply -k ./todo-project/
