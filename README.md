# pubsub-sample

## Requirement
- Docker Desktop
- VSCode with the extension of `Dev Containers`

The maintainer uses GitHub Codespaces

## How to use
```
# create topics and subscriptions
$ make setup

# start subscription
$ make subscribe

# on another terminal, publish a message
$ make msg=${your message} publish
```