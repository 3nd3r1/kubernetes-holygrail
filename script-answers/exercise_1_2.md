```bash
3nd3r1@mypc$ cd ./todo-project/backend
3nd3r1@mypc$ docker build -f ./build/Dockerfile -t 3nd3r1/todo-backend .
3nd3r1@mypc$ docker push 3nd3r1/todo-backend:latest
3nd3r1@mypc$ kubectl create deployment todo-backend --image=3nd3r1/todo-backend
```
