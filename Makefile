.PHONY: topic publish run-publisher run-subscriber

topic:
	chmod 777 .devcontainer/pubsub/topic.sh
	.devcontainer/pubsub/topic.sh

publish:
	curl http://localhost:8080/publish

run-publisher:
	go run main.go publisher

run-subscriber:
	go run main.go subscriber