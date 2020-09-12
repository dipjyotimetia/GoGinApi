## Kubernetes deployment descriptions
### Adding Secrets
```shell script  
kubectl create -f config/kubernetes/postgres-secret.yaml  
kubectl describe secrets postgres-secret  
```
### Adding configs
```shell script
kubectl apply -f config/kubernetes/postgres-configmap.yaml
```
### Create PersistentVolume
```shell script
kubectl apply -f config/kubernetes/postgres-db-pv.yaml
```
### Create a PersistentVolumeClaim to request a PersistentVolume storage
```shell script
kubectl apply -f config/kubernetes/postgres-db-pvc.yaml
```
### Deploy postgres db
```shell script
kubectl apply -f config/kubernetes/postgres-db-deployment.yaml
```    
### Start postgres db service
```shell script
kubectl apply -f config/kubernetes/postgres-db-service.yaml
```
### Deploy api
```shell script
kubectl apply -f config/kubernetes/app-postgres-deployment.yaml
```
### Start api service
```shell script
kubectl apply -f config/kubernetes/app-postgres-service.yaml
```
### Get pods
```shell script
kubectl get pods
```    
### Get Services
```shell script
kubectl get services
kubectl get services fullstack-app-postgres
```
###Display information about the Deployment:
```shell script
kubectl get deployments fullstack-app-postgres
kubectl describe deployments fullstack-app-postgres
```
###Create a Service object that exposes the deployment:
```shell script
kubectl expose deployment fullstack-app-postgres --type=LoadBalancer --name=goginapi
```
###Display information about the Service:
```shell script
kubectl get services goginapi
```