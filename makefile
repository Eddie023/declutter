build:
	docker build -t declutter:0.0.1 . 

run:
	docker run decluter:0.0.1 

test:
	docker run -t \
	--volume $$(pwd):/go/src/github.com/eddie023/declutter \
	--rm  \
	$$(docker build --quiet --file test.Dockerfile . ) 

lint: 
	docker run --rm \
	--volume $$(pwd):/src \
	--volume ~/.cache:/root./.cache \
	$$(docker build --quiet --file lint.Dockerfile .) \
	golangci-lint run 

release:
	rm -rf dist/ && goreleaser build --single-target --skip-validate 