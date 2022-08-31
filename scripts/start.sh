BASEDIR=$(dirname "$0")

cd /app

source $BASEDIR/internal/config.sh
source $BASEDIR/internal/set-config.sh

aws --endpoint-url $AWS_ENDPOINT lambda delete-function --function-name aws-sqs-lambda-golang

aws --endpoint-url $AWS_ENDPOINT sqs create-queue --queue-name aws-sqs-lambda-golang-queue > /dev/null
aws --endpoint-url $AWS_ENDPOINT lambda create-function --function-name aws-sqs-lambda-golang --runtime go1.x --handler main --zip-file fileb://app.zip --role aws:iam::123456789012:role/execution_role
rm -rf app.zip main
aws --endpoint-url $AWS_ENDPOINT lambda create-event-source-mapping --function-name aws-sqs-lambda-golang --batch-size 5 --maximum-batching-window-in-seconds 60 --event-source-arn arn:aws:sqs:us-east-1:000000000000:aws-sqs-lambda-golang-queue
