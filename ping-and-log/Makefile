build-docker:
	docker build -t 3nd3r1/log-output:latest -f ./log-output/build/Dockerfile ./log-output
	docker build -t 3nd3r1/ping-pong:latest -f ./ping-pong/build/Dockerfile ./ping-pong
	docker push 3nd3r1/log-output:latest
	docker push 3nd3r1/ping-pong:latest

install-linkerd:
	# Install linkerd
	linkerd install --crds | kubectl apply -f - || true
	linkerd install | kubectl apply -f - || true
	linkerd viz install | kubectl apply -f - || true

deploy: install-linkerd
	# Deploy ping-and-log
	echo "Deploying ping-and-log"
	kubectl create namespace ping-and-log --dry-run=client -o yaml | kubectl apply -f -
	kubectl config set-context --current --namespace=ping-and-log
	kubectl apply -f ./ping-pong/manifests/secrets/secret.yaml
	kubectl apply -k ./ping-pong/
	kubectl apply -k ./log-output/

deploy-knative: install-linkerd
	# Deploy ping-and-log
	echo "Deploying ping-and-log with Knative"
	kubectl create namespace ping-and-log --dry-run=client -o yaml | kubectl apply -f -
	kubectl config set-context --current --namespace=ping-and-log
	kubectl apply -f ./ping-pong/manifests/secrets/secret.yaml
	kubectl apply -k ./ping-pong/knative
	kubectl apply -k ./log-output/knative

uninstall:
	kubectl delete namespace ping-and-log --ignore-not-found
