build-docker:
	docker build -t 3nd3r1/website-copyer:latest -f ./website-copyer/Dockerfile ./website-copyer
	docker push 3nd3r1/website-copyer:latest

deploy:
	kubectl apply -f ./manifests/

deploy-example:
	kubectl create namespace dummysite --dry-run=client -o yaml | kubectl apply -f -
	kubectl apply -f ./examples/

uninstall:
	kubectl delete namespace system --ignore-not-found
	kubectl delete namespace dummysite --ignore-not-found
