```bash
3nd3r1@mypc$ k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
3nd3r1@mypc$ cd ./todo-project
3nd3r1@mypc$ kubectl apply -f ./backend/manifests/deployment.yaml
3nd3r1@mypc$ kubectl apply -f ./backend/manifests/service.yaml
```
