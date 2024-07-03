```bash
3nd3r1@mypc$ cd ./log-output
3nd3r1@mypc$ docker build -f ./build/Dockerfile -t 3nd3r1/log-output .
3nd3r1@mypc$ docker push 3nd3r1/log-output:latest
3nd3r1@mypc$ k3d cluster create -a 2
3nd3r1@mypc$ kubectl create deployment log-output --image=3nd3r1/log-output
3nd3r1@mypc$ kubectl logs log-output-b44c49b69-gbmbn -f
```
