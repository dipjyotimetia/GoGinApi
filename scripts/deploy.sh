#!/bin/bash
@echo "Kubernetes up"
	kubectl create -f deployments/kubernetes/postgres-secret.yaml
	kubectl describe secrets postgres-secret
	kubectl apply -f deployments/kubernetes/postgres-configmap.yaml
	kubectl apply -f deployments/kubernetes/postgres-db-pv.yaml
  kubectl apply -f deployments/kubernetes/postgres-db-pvc.yaml
  kubectl apply -f deployments/kubernetes/postgres-db-deployment.yaml
  kubectl apply -f deployments/kubernetes/postgres-db-service.yaml

@echo "Wait for postgres db to up"
sleep 1.5

@echo "helm chart deploy services"
  cd config
	helm lint
	helm install goginapi helmchart
	kubectl expose deployment goginapi-helmchart --type=LoadBalancer --name=goginapi


