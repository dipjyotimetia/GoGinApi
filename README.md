## Go Gin Api

### MockGen

https://github.com/IsraelAdura/go-postgres-jwt-react-starter
```bash
mockgen -source=repository/user-repository.go -destination=mocks/user-mock/mock_repository.go

```
### Kubernetes deployment
```shell script
minikube start

kubectl create -f postgres-secret.yaml
kubectl get secrets

kubectl describe secrets  postgres-secret

kubectl apply -f postgres-db-pv.yaml
kubectl apply -f postgres-db-pvc.yaml
kubectl apply -f postgres-db-deployment.yaml
kubectl apply -f postgres-db-service.yaml

kubectl get pods

kubectl describe pod <pod_name>
kubectl logs <pod_name>

docker build -t <app-name> .
docker tag <image-name> <dockerhub-username>/<repository-name>:<tag-name>

docker push <dockerhub-username>/<repository-name>:<tag_name>

kubectl apply -f app-postgres-deployment.yaml
kubectl apply -f app-postgres-service.yaml

kubectl get pods

kubectl describe pod <pod-name>
kubectl logs <pod-name>

kubectl get services

minikube service fullstack-app-postgres --url

minikube stop
```
