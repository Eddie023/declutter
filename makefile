run: 
	go run cmd/declutter/main.go /Users/eddie/dev/go/src/declutter/testFiles

test:
	go test internal/** -v

reset-testFiles:
	find testFiles -maxdepth 2 -type f -exec mv {} ./testFiles \;

reset-run: reset-testFiles run 