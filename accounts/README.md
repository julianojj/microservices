# Accounts

## Boas práticas
Utilizamos o [semver](https://semver.org/), para padronizar os commits:
```
feat: commit message
fix: commit message
docs: commit message
style: commit message
refactor: commit message
test: commit message
chore: commit message
perf: commit message
ci: commit message
build: commit message
temp: commit message
```

## Principais tecnologias
```
go: v1.19.2
github.com/google/uuid: v1.3.0
github.com/stretchr/testify: v1.8.1
github.com/golang-jwt/jwt/v4: v4.4.3
```

## Comandos utilizados
```
go test --cover ./...
go test -coverprofile=cover.out ./...
go tool cover -html=cover.out
```
