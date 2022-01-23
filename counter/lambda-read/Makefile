create-table:
	aws dynamodb create-table \
		--table-name cloud-resume-challenge-counter \
		--attribute-definitions \
			AttributeName=PK,AttributeType=S \
		--key-schema \
			AttributeName=PK,KeyType=HASH \
		--billing-mode PAY_PER_REQUEST \
		--endpoint-url http://localhost:8000

test:
	LOCAL=true go test ./...
