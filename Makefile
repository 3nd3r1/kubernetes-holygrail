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
