## Go Gin Api

### Architecture
<img src="https://github.com/dipjyotimetia/GoGinApi/blob/master/docs/architecture/goginapi.png" width="500"> 

### To Generate Mock
```bash
mockgen -source=repository/user-repository.go -destination=mocks/user-mock/mock_repository.go
```

### Run in docker
```bash
make compose | docker-compose up -d
```

### Run in local kubernetes cluster
```bash
deploy.sh
cleandeploy.sh
```

https://timothyagustian.medium.com/simulating-customized-chaos-in-golang-using-toxiproxy-b913584d88a7