#/bin/bash

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
