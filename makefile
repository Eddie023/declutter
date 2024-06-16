VERSION := 0.0.1 

build:
	docker build \
	-t declutter:${VERSION} \
	--build-arg VERSION=${VERSION} \
	. 

run:
	docker run decluter:${VERSION} 

test:
	docker run -t \
	--volume $$(pwd):/go/src/github.com/eddie023/declutter \
	--rm  \
	$$(docker build --quiet --file test.Dockerfile . ) 

reset-testFiles:
	find testFiles -maxdepth 2 -type f -exec mv {} ./testFiles \;

lint: 
	docker run --rm \
	--volume $$(pwd):/src \
	--volume ~/.cache:/root./.cache \
	$$(docker build --quiet --file lint.Dockerfile .) 

release:
	rm -rf dist/ && goreleaser build --single-target --skip-validate 