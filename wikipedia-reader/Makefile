deploy:
	echo "Deploying wikipedia-reader"
	kubectl create namespace wikipedia-reader --dry-run=client -o yaml | kubectl apply -f -
	kubectl config set-context --current --namespace=wikipedia-reader
	kubectl apply -f ./manifests/

uninstall:
	echo "Uninstalling wikipedia-reader"
	kubectl delete namespace wikipedia-reader
