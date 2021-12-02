APP_NAME = christmastree
BASE_GO_IMAGE = golang:1.17.3-alpine3.14
BASE_TARGET_IMAGE = alpine:3.14

REGISTRY ?= localhost:5000
IMAGE_DEV = $(APP_NAME)-dev
IMAGE_PROD = $(APP_NAME)-application


DOCKERFILE_DEV = .docker/dev/Dockerfile
DOCKERFILE_PROD = ./Dockerfile.dev


CGO_ENABLED = 0 # statically linked = 0
GOOS=linux
GOARCH=arm
GOARM=5

.DEFAULT_GOAL = go-build
PID = /tmp/serving.pid
DEVELOPER_UID     ?= $(shell id -u)
#-----------------------------------------------------------------------------------------------------------------------
ARG := $(word 2, $(MAKECMDGOALS))
%:
	@:
#-----------------------------------------------------------------------------------------------------------------------
#-----------------------------------------------------------------------------------------------------------------------

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

build-image-dev: ## Build dev image
	@docker build 											\
	    -t $(REGISTRY)/$(IMAGE_DEV)            			    \
		--build-arg BASE_GO_IMAGE=$(BASE_GO_IMAGE)          \
		--build-arg DEVELOPER_UID=$(DEVELOPER_UID)          \
		--build-arg APP_NAME=${APP_NAME}					\
		-f $(DOCKERFILE_DEV)								\
		.

# build-image-prod: ## Build prod image
# 	@docker build 											\
# 	    -t $(REGISTRY)/$(IMAGE_PROD)						\
# 		--build-arg BASE_GO_IMAGE=$(BASE_GO_IMAGE)			\
# 		--build-arg BASE_TARGET_IMAGE=$(BASE_TARGET_IMAGE)	\
# 		--build-arg CGO_ENABLED=${CGO_ENABLED}				\
# 		--build-arg GOOS=${GOOS}							\
# 		--build-arg GOARCH=${GOARCH}						\
# 		--build-arg GOARM=${GOARM}							\
# 		--build-arg APP_NAME=${APP_NAME}					\
# 		-f $(DOCKERFILE_PROD)								\
# 		.

run-image: ## Run prod image
	@docker run ${REGISTRY}/$(IMAGE_PROD)

up: ## Start application dev container
	@cd .docker && docker-compose up -d
down: ## Remove application dev container
	@cd .docker && docker-compose down
console: ## Enter application dev container
	@docker exec -it -u developer $(IMAGE_DEV) bash
console-root: ## Enter application dev container
	@docker exec -it -u root $(IMAGE_DEV) bash

go-build: ## Build dev application (go build)	
	@env CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -o bin/${APP_NAME} ./
clean: ## Clean bin/
	@rm -rf bin/${APP_NAME}
