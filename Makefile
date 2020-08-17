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
	go run server.go

.PHONY:compose
compose:
	@echo build docker image
	docker-compose down
	docker-compose up --remove-orphans

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
	kubectl describe secrets  postgres-secret
	kubectl apply -f config/kubernetes/postgres-db-pv.yaml
    kubectl apply -f config/kubernetes/postgres-db-pvc.yaml
    kubectl apply -f config/kubernetes/postgres-db-deployment.yaml
    kubectl apply -f config/kubernetes/postgres-db-service.yaml
    kubectl apply -f config/kubernetes/app-postgres-deployment.yaml
    kubectl apply -f config/kubernetes/app-postgres-service.yaml