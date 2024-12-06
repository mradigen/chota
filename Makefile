deps:
	go mod tidy

run: deps
	go run cmd/app.go

build: deps
	go build -o short cmd/app.go

test:
	go test tests/app_test.go -v

docker:
	docker build -t short .

kubernetes:
	# envsubst applies environment variables like DOCKER_REGISTRY
	kubectl apply -f deploy/kubernetes.yml
	
#disallow any parallelism
.NOTPARALLEL:
