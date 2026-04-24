.PHONY: topic publish

topic:
	chmod 777 .devcontainer/pubsub/topic.sh
	.devcontainer/pubsub/topic.sh

publish:
	curl http://localhost:8080/publish

run-publisher:
	go run main.go publisher