docker-build:
	docker build -t 3nd3r1/todo-project-backend:latest -f ./todo-project/backend/build/Dockerfile ./todo-project/backend
	docker build -t 3nd3r1/todo-project-frontend:latest -f ./todo-project/frontend/build/Dockerfile ./todo-project/frontend
	docker build -t 3nd3r1/todo-project-imagenator:latest -f ./todo-project/imagenator/build/Dockerfile ./todo-project/imagenator
	docker build -t 3nd3r1/todo-project-backup-agent:latest -f ./todo-project/backup-agent/build/Dockerfile ./todo-project/backup-agent
	docker build -t 3nd3r1/todo-project-broadcaster:latest -f ./todo-project/broadcaster/build/Dockerfile ./todo-project/broadcaster
	docker push 3nd3r1/todo-project-backend:latest
	docker push 3nd3r1/todo-project-frontend:latest
	docker push 3nd3r1/todo-project-imagenator:latest
	docker push 3nd3r1/todo-project-backup-agent:latest
	docker push 3nd3r1/todo-project-broadcaster:latest

install-prometheus:
	kubectl create namespace prometheus --dry-run=client -o yaml | kubectl apply -f -
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add stable https://charts.helm.sh/stable
	helm upgrade --install kube-prometheus-stack prometheus-community/kube-prometheus-stack --namespace prometheus

install-argo-rollouts:
	kubectl create namespace argo-rollouts --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml

install-argocd:
	kubectl create namespace argocd --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

install-linkerd:
	linkerd install --crds | kubectl apply -f - || true
	linkerd install | kubectl apply -f - || true
	linkerd viz install | kubectl apply -f - || true

install: install-prometheus install-argo-rollouts install-argocd install-linkerd

deploy-staging: install
	echo "Deploying todo-project-staging"
	kubectl create namespace staging --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n staging -f ./manifests/secrets/secret.staging.yaml
	helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false staging-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=staging
	kubectl apply -n argocd -f ./manifests/application.staging.yaml

	echo "Waiting for argocd server:"
	kubectl -n argocd wait --for=condition=ready pod -l app.kubernetes.io/name=argocd-server --timeout=300s
	echo "Argocd admin password:"
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

deploy-prod: install
	echo "Deploying todo-project-prod"
	kubectl create namespace production --dry-run=client -o yaml | kubectl apply -f -
	sops --decrypt ./manifests/secrets/secret.enc.yaml | kubectl apply -n production -f -
	helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false prod-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=production
	kubectl apply -n argocd -f ./manifests/application.prod.yaml

	echo "Waiting for argocd server:"
	kubectl -n argocd wait --for=condition=ready pod -l app.kubernetes.io/name=argocd-server --timeout=300s
	echo "Argocd admin password:"
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

uninstall:
	kubectl delete namespace staging --ignore-not-found
	kubectl delete namespace production --ignore-not-found
