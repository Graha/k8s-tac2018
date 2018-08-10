# k8s-tac2018

## Deployment
$ kubectl apply -f deployment.yaml
$ kubectl describe deployment go-app
$ kubectl get pods -l app=go-app


## Manual Scale
$ kubectl scale --replicas=3 deployment/go-app

## Auto Scale
$ kubectl autoscale deployment/go-app --min=3 --max=5 --cpu-percent=20
$ ab -n 10000 -c 100 localhost:30000/

## Usage
$ curl localhost:3000/
go-app-7df97fd879-hmnzr : Welcome to Ping/Pong (v1)...

$ curl localhost:3000/ping
go-app-7df97fd879-hmnzr : pong

$ curl localhost:3000/pong
go-app-7df97fd879-hmnzr : ping