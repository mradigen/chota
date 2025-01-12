deps:
	go mod tidy

run: deps
	go run cmd/app.go

build: deps
	go build -o chota cmd/app.go

test:
	go test tests/app_test.go -v

docker:
	docker build -t chota .

kubernetes:
	# envsubst applies environment variables like DOCKER_REGISTRY
	kubectl apply -f deploy/kubernetes.yml
	
#disallow any parallelism
.NOTPARALLEL:
