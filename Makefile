create-cluster:
	k3d cluster delete
	k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
	docker exec k3d-k3s-default-agent-0 mkdir /tmp/kube
	docker exec k3d-k3s-default-agent-0 chmod -R 0777 /tmp/kube

create-cluster-knative:
	k3d cluster delete
	k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2 --k3s-arg "--disable=traefik@server:0"
	docker exec k3d-k3s-default-agent-0 mkdir /tmp/kube
	docker exec k3d-k3s-default-agent-0 chmod -R 0777 /tmp/kube
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.16.0/serving-crds.yaml
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.16.0/serving-core.yaml
	kubectl apply -f https://github.com/knative/net-kourier/releases/download/knative-v1.16.0/kourier.yaml
	kubectl patch configmap/config-network \
	  --namespace knative-serving \
	  --type merge \
	  --patch '{"data":{"ingress-class":"kourier.ingress.networking.knative.dev"}}'
	kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.16.0/serving-default-domain.yaml

build-applications:
	docker build -t 3nd3r1/log-output:latest -f ./log-output/build/Dockerfile ./log-output
	docker build -t 3nd3r1/ping-pong:latest -f ./ping-pong/build/Dockerfile ./ping-pong
	docker push 3nd3r1/log-output:latest
	docker push 3nd3r1/ping-pong:latest

build-todo-project:
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

deploy-todo-project:
	export SOPS_AGE_KEY_FILE=./key.txt

	# Deploy prometheus
	kubectl create namespace prometheus --dry-run=client -o yaml | kubectl apply -f -
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo add stable https://charts.helm.sh/stable
	helm install prometheus-community/kube-prometheus-stack --generate-name --namespace prometheus

	# Deploy argo-rollouts
	kubectl create namespace argo-rollouts --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml

	# Deploy argocd
	kubectl create namespace argocd --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

	# Install linkerd
	linkerd install --crds | kubectl apply -f -
	linkerd install | kubectl apply -f -
	linkerd viz install | kubectl apply -f -

	# Setup todo-project
	echo "Deploying todo-project"
	kubectl create namespace production --dry-run=client -o yaml | kubectl apply -f -
	kubectl create namespace staging --dry-run=client -o yaml | kubectl apply -f -
	sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -n production -f -
	sops --decrypt ./todo-project/manifests/secrets/secret.enc.yaml | kubectl apply -n staging -f -
	helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false prod-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=production
	helm upgrade --install --set commonLabels.project=todo-project --set auth.enabled=false staging-todo-project-nats oci://registry-1.docker.io/bitnamicharts/nats --namespace=staging
	kubectl apply -n argocd -f ./todo-project/manifests/application.yaml

	echo "Waiting for argocd server:"
	kubectl -n argocd wait --for=condition=ready pod -l app.kubernetes.io/name=argocd-server --timeout=300s
	echo "Argocd admin password:"
	kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d

deploy-applications:
	# Install linkerd
	linkerd install --crds | kubectl apply -f -
	linkerd install | kubectl apply -f -
	linkerd viz install | kubectl apply -f -

	# Deploy ping-pong and log-output
	echo "Deploying ping-pong and log-output"
	kubectl create namespace applications --dry-run=client -o yaml | kubectl apply -f -
	kubectl config set-context --current --namespace=applications
	sops --decrypt ./ping-pong/manifests/secrets/secret.enc.yaml | kubectl apply -f -
	kubectl apply -k ./ping-pong/
	kubectl apply -k ./log-output/
