run:
	go run cmd/app.go

build:
	go build -o short cmd/app.go

test:
	go test tests/app_test.go -v

docker:
	docker build -t ${DOCKER_REGISTRY}short .

kubernetes:
	# envsubst applies environment variables like DOCKER_REGISTRY
	cat deploy/kubernetes.yml | envsubst | kubectl apply -f -
	
#disallow any parallelism
.NOTPARALLEL:
