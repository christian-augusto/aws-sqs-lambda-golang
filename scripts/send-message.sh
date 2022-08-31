BASEDIR=$(dirname "$0")
SQS_QUEUE_NAME="aws-sqs-lambda-golang-queue"
MESSAGE_BODY="{\"phone\": \"phone number\",\"message\": \"hello world 2\"}"

source $BASEDIR/internal/config.sh

aws --endpoint-url $AWS_ENDPOINT sqs send-message --queue-url $AWS_ENDPOINT/000000000000/$SQS_QUEUE_NAME --message-body "$MESSAGE_BODY" > /dev/null
