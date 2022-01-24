.PHONY: build

build:
	sam build

deploy: build
	sam deploy --no-confirm-changeset --no-fail-on-empty-changeset --force-upload

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
	cd lambda-incr && LOCAL=true go test ./...
	cd lambda-read && LOCAL=true go test ./...
