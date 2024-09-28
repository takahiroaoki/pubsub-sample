setup:
	chmod 777 setup.sh
	./setup.sh
	cd ./subscriber && go mod tidy
	cd ./publisher && go mod tidy

subscribe:
	cd ./subscriber && go run main.go

publish:
	cd ./publisher && go run main.go -msg=$(msg)