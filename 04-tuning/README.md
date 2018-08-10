# k8s-tac2018

## Quota
$ kubectl apply -f deployment-quota.yaml

### heapster available
$ kubectl top pod <pod-name>  

### if not
$ kubectl describe pod <pod-name> 

## Quality of Service
$ kubectl create namespace qos-try
$ kubectl create -f deployment-qos.yaml --namespace=qos-try
$ kubectl describe pod <pod-name>  --namespace=qos-try  | grep QoS  

## Health
$ kubectl apply -f deployment-health.yaml
$ kubectl describe deployment go-app

