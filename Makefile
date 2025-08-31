
CONSUMER_KEY?=
CONSUMER_SECRET?=
ACCESS_TOKEN?=
ACCESS_TOKEN_SECRET?=

build:
	go mod download && go build .

test:
	mkdir -p coverage && go test ./... --cover -coverprofile coverage/coverage.out

create-tweet:
	CONSUMER_KEY=${CONSUMER_KEY} CONSUMER_SECRET=${CONSUMER_SECRET} ACCESS_TOKEN=${ACCESS_TOKEN} ACCESS_TOKEN_SECRET=${ACCESS_TOKEN_SECRET}  go run ./compose

swagger:
	echo "Running Swagger UI on port 8080"
	@docker run -p 8080:8080 -e SWAGGER_JSON=/app/openapi.yaml -v ${PWD}/docs:/app swaggerapi/swagger-ui

.PHONY: build test create-tweet swagger
