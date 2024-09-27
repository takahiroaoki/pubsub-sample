#/bin/bash

set -em

# start emulator in background
gcloud beta emulators pubsub start --project=$PUBSUB_PROJECT_ID --host-port=$PUBSUB_EMULATOR_HOST --quiet &

# wait until the emulator has started
while ! nc -z localhost 8085; do
  sleep 0.1
done

# create topics
ENDPOINT="${PUBSUB_EMULATOR_HOST}/v1/projects/${PUBSUB_PROJECT_ID}"
curl -S -XPUT "${ENDPOINT}/topics/${PUBSUB_TOPIC_ID}"
curl -S -XPUT "${ENDPOINT}/topics/${PUBSUB_DEAD_LETTER_TOPIC_ID}"

# create subscriptions
curl -S -XPUT "${ENDPOINT}/subscriptions/${PUBSUB_SUBSCRIPTION_ID}" \
  -H "content-type: application/json" \
  -d '{
    "topic": "projects/'${PUBSUB_PROJECT_ID}'/topics/'${PUBSUB_TOPIC_ID}'",
    "deadLetterPolicy": {
      "deadLetterTopic": "projects/'${PUBSUB_PROJECT_ID}'/topics/'${PUBSUB_DEAD_LETTER_TOPIC_ID}'",
      "maxDeliveryAttempts": 5
    }
  }'
curl -S -XPUT "${ENDPOINT}/subscriptions/${PUBSUB_DEAD_LETTER_SUBSCRIPTION_ID}" \
  -H "content-type: application/json" \
  -d '{
    "topic": "projects/'${PUBSUB_PROJECT_ID}'/topics/'${PUBSUB_DEAD_LETTER_TOPIC_ID}'"
  }'

# move the emulator process foreground
fg %1
