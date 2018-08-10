# k8s-tac2018

## Deployment
$ kubectl apply -f deployment.yaml
$ kubectl describe deployment go-app
$ kubectl get pods -l app=go-app,track=stable

## Canary Deployment
$ kubectl apply -f deployment-canary.yaml
$ kubectl describe deployment go-app-canary
$ kubectl get pods -l app=go-app,track=canary

## Usage
$ curl localhost:3000/
go-app-7df97fd879-hmnzr : Welcome to Ping/Pong (v1)...

$ curl localhost:3000/ping
go-app-7df97fd879-hmnzr : pong

$ curl localhost:3000/pong
go-app-7df97fd879-hmnzr : ping