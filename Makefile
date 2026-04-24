setup:
	chmod 777 .devcontainer/pubsub/setup.sh
	.devcontainer/pubsub/setup.sh

subscribe:
	cd ./subscriber && go run main.go

publish:
	cd ./publisher && go run main.go -msg=$(msg)