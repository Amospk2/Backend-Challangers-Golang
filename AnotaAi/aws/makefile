build:
	@echo "Creating SQS queue..."
	aws --endpoint-url http://localhost:4566 sqs create-queue --queue-name my-queue

	@echo "Creating SNS topic..."
	aws --endpoint-url http://localhost:4566 sns create-topic --name my-topic

	@echo "Subscribing SQS queue to SNS topic..."
	aws --endpoint-url http://localhost:4566 sns subscribe \
		--topic-arn "arn:aws:sns:us-east-1:000000000000:my-topic"  \
		--protocol sqs \
		--notification-endpoint "arn:aws:sqs:us-east-1:000000000000:my-queue"


	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main -ldflags '-w' main.go && zip archive.zip main
	
	aws --endpoint-url http://localhost:4566 lambda create-function --function-name my-lambda \
		--zip-file fileb://archive.zip --handler main --runtime go1.x \
		--role arn:aws:iam::123456789012:role/my-lambda

	@echo "Creating event source mapping for Lambda function..."
	aws --endpoint-url http://localhost:4566 lambda create-event-source-mapping \
		--function-name my-lambda \
		--batch-size 5 \
		--maximum-batching-window-in-seconds 60 \
		--event-source-arn arn:aws:sqs:us-east-1:000000000000:my-queue

	@echo "Creating S3 bucket..."
	aws --endpoint-url http://localhost:4566 s3 mb s3://anotaai-catalog-marketplace

	@echo "Build completed."






create_lambda:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main -ldflags '-w' main.go && zip archive.zip main
	
	aws --endpoint-url http://localhost:4566 lambda create-function --function-name my-lambda \
		--zip-file fileb://archive.zip --handler main --runtime go1.x \
		--role arn:aws:iam::123456789012:role/my-lambda

	

