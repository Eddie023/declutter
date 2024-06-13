build:
	docker build -t declutter:0.0.1 . 

run:
	docker run decluter:0.0.1 

test:
	docker run \
	--volume $$(pwd):/go/src/github.com/eddie023/declutter \
	--rm --interactive --tty \
	$$(docker build --quiet --file test.Dockerfile . ) 

lint: 
	docker run \
	--rm --interactive --tty \
	$$(docker build --quiet --file lint.Dockerfile . ) > /dev/null 2>&1 \
	&& golangci-lint run 

release:
	rm -rf dist/ && goreleaser build --single-target --skip-validate 