#!/usr/bin/env bash
set -eo pipefail

ENDPOINT="http://localhost:4566"
QUEUE_NAME="task-queue.fifo"

# disables the use of a pager
export AWS_PAGER=""

# create queue
aws --endpoint ${ENDPOINT} \
    sqs create-queue \
    --queue-name ${QUEUE_NAME} \
    --attributes '{"FifoQueue": "true", "ContentBasedDeduplication":"true"}'

COUNTER=1
for ((n = 0; n < 100; n++)); do
    aws --endpoint ${ENDPOINT} \
        sqs send-message \
        --queue-url "${ENDPOINT}/queue/${QUEUE_NAME}" \
        --message-body "{'msg': 'hello-${COUNTER}'}" \
        --message-group-id "a-string"

    COUNTER=$(($COUNTER + 1))
done
