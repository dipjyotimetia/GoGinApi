BUILDPATH = ${CURDIR}

EXENAME = main

makedir:
	@if [ ! -d ${BUILDPATH}/bin ] ; then mkdir -p ${BUILDPATH}/bin ; fi
	@if [ ! -d ${BUILDPATH}/pkg ] ; then mkdir -p ${BUILDPATH}/pkg ; fi
build:
	@echo "Start building..."
	go build -o bin/main server.go

clean:
	@echo "Cleaning..."
	@rm -rf ${BUILDPATH}/bin/${EXENAME}
	@rm -rf ${BUILDPATH}/pkg

run:
	go run server.go

compose:
	@echo build docker image
	docker-compose up --remove-orphans

prune:
	@echo clean all that is not actively used
	docker system prune -af

all: makedir build

update:
	@echo "Updating..."
	go get -t -u ./...
