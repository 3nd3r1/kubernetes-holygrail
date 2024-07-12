```bash
3nd3r1@mypc$ docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/kube
3nd3r1@mypc$ docker exec k3d-k3s-default-agent-0 chmod -R 0777 /tmp/kube
3nd3r1@mypc$ ./build.sh
3nd3r1@mypc$ ./delete.sh
3nd3r1@mypc$ ./deploy.sh
```
