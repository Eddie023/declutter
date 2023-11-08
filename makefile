run: 
	go run cmd/declutter/main.go -s -p /Users/eddie/dev/go/src/declutter/testFiles

reset-testFiles:
	find testFiles -maxdepth 2 -type f -exec mv {} ./testFiles \;

reset-run: reset-testFiles run 

build:
	rm -rf dist/ && goreleaser build --single-target --skip-validate 