# devops-journey

[![Go App CI/CD](https://github.com/shambhand/devops-journey/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/shambhand/devops-journey/actions/workflows/ci-cd.yaml)

## Week 1: Containerization
### run app
``` sh
go run ./src/main.go
```

### run test
```sh
go test ./... 
```

### docker-compose
```sh
docker-compose up -d
```

## Week 2: Tools
* go
* docker and docker-compose
* kubectl
* helm
* minikube

## Week 3: Kubernetes and Helm
``` sh
helm template <release-name> <chart-dir>
helm upgrade --install <release-name> <chart.tgz>
```
## Week 4: CI with Github Actions
* Run Tests -> Docker build and push -> helm package and upload

## Week 5: CD with Github Actions
* Setup self hosted runner on local machine
* Deploy to local minikube cluster with the self hosted runner
  * Deploy: DEV -> QA -> Prod

