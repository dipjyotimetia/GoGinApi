BUILDPATH = ${CURDIR}

EXENAME = main

.PHONY:makedir
makedir:
	@if [ ! -d ${BUILDPATH}/bin ] ; then mkdir -p ${BUILDPATH}/bin ; fi
	@if [ ! -d ${BUILDPATH}/pkg ] ; then mkdir -p ${BUILDPATH}/pkg ; fi

.PHONY:build
build:
	@echo "Start building..."
	go build -o bin/main server.go

.PHONY:clean
clean:
	@echo "Cleaning..."
	@rm -rf ${BUILDPATH}/bin/${EXENAME}
	@rm -rf ${BUILDPATH}/pkg

.PHONY:run
run:
	go run cmd/server.go

.PHONY:compose
compose:
	@echo build docker image
	docker-compose down --remove-orphans
	docker-compose up

.PHONY:prune
prune:
	@echo clean all that is not actively used
	docker system prune -af

all: makedir build

.PHONY:update
update:
	@echo "Updating..."
	go get -t -u ./...

.PHONY:kube-up
kube-up:
	@echo "Kubernetes up"
	kubectl create -f config/kubernetes/postgres-secret.yaml
	kubectl describe secrets postgres-secret
	kubectl apply -f config/kubernetes/postgres-configmap.yaml
	kubectl apply -f config/kubernetes/postgres-db-pv.yaml
    kubectl apply -f config/kubernetes/postgres-db-pvc.yaml
    kubectl apply -f config/kubernetes/postgres-db-deployment.yaml
    kubectl apply -f config/kubernetes/postgres-db-service.yaml
    kubectl apply -f config/kubernetes/app-postgres-deployment.yaml
    kubectl apply -f config/kubernetes/app-postgres-service.yaml
    kubectl get pods

    kubectl get services
    kubectl get services fullstack-app-postgres

	#Display information about the Deployment:
 	kubectl get deployments fullstack-app-postgres
 	kubectl describe deployments fullstack-app-postgres

	#Create a Service object that exposes the deployment:
    kubectl expose deployment fullstack-app-postgres --type=LoadBalancer --name=goginapi
    #Display information about the Service:
    kubectl get services goginapi

.PHONY:kube-down
kube-down:
	@echo "Kubernetes down"
	kubectl delete services goginapi
	kubectl delete deployment fullstack-app-postgres
	kubectl delete configmap postgres-config
	kubectl delete persistentvolumeclaim postgres-pv-claim
    kubectl delete persistentvolume postgres-pv-volume
	kubectl delete secret postgres-secret
	kubectl delete deployment --all
	kubectl delete pods --all

.PHONY:kube-clean
kube-clean:
	kubectl delete deployments --all
    kubectl delete services --all
    kubectl delete pods --all
    kubectl delete daemonset --all

.PHONY:kube-dash-token
kube-dash-token:
	@echo "Kubernetes dashboard token generate"
	kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
   #kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | sls admin-user | ForEach-Object { $_ -Split '\s+' } | Select -First 1)
   #http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

.PHONY:kube-dashboard
kube-dashboard:
	@echo "Kubernetes dashboard"
	kubectl proxy