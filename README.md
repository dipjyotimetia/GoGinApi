## Go Gin Api

### To Generate Mock
```bash
mockgen -source=repository/user-repository.go -destination=mocks/user-mock/mock_repository.go
```

### Run in docker
```bash
make compose
```

docker-compose -f docker-compose-infra up -d
