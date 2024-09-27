#/bin/bash

set -em

# start emulator in background
gcloud beta emulators pubsub start --project=$PROJECT_ID --host-port=$EMULATOR_HOST --quiet &

# wait until the emulator has started
while ! nc -z localhost 8085; do
  sleep 0.1
done

# create topics
ENDPOINT="${EMULATOR_HOST}/v1/projects/${PROJECT_ID}"
curl -S -XPUT "${ENDPOINT}/topics/${TOPIC_ID}"
curl -S -XPUT "${ENDPOINT}/topics/${DEAD_LETTER_TOPIC_ID}"

# create subscriptions
curl -S -XPUT "${ENDPOINT}/subscriptions/${SUBSCRIPTION_ID}" \
  -H "content-type: application/json" \
  -d '{
    "topic": "projects/'${PROJECT_ID}'/topics/'${TOPIC_ID}'",
    "deadLetterPolicy": {
      "deadLetterTopic": "projects/'${PROJECT_ID}'/topics/'${DEAD_LETTER_TOPIC_ID}'",
      "maxDeliveryAttempts": 5
    }
  }'
curl -S -XPUT "${ENDPOINT}/subscriptions/${DEAD_LETTER_SUBSCRIPTION_ID}" \
  -H "content-type: application/json" \
  -d '{
    "topic": "projects/'${PROJECT_ID}'/topics/'${DEAD_LETTER_TOPIC_ID}'"
  }'

# move the emulator process foreground
fg %1
