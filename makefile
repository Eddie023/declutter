run: 
	go run cmd/declutter/main.go ${path}

test:
	go test internal/** -v