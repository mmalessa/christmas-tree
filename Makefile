include .env

APP_NAME = christmastree
BASE_GO_IMAGE = golang:1.17.3-alpine3.14
BASE_TARGET_IMAGE = alpine:3.14

REGISTRY ?= localhost:5000
IMAGE_DEV = $(APP_NAME)-dev

DOCKERFILE_DEV = .docker/images/dev/Dockerfile

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
## HELP
help: ## Outputs this help screen
	@grep -hE '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

## Host
config: ## Build dev image
	@docker build 											\
	    -t $(REGISTRY)/$(IMAGE_DEV)            			    \
		--build-arg BASE_GO_IMAGE=$(BASE_GO_IMAGE)          \
		--build-arg DEVELOPER_UID=$(DEVELOPER_UID)          \
		--build-arg APP_NAME=${APP_NAME}					\
		-f $(DOCKERFILE_DEV)								\
		.

up: ## Start application dev container
	@cd .docker && \
	APP_NAME=$(APP_NAME) \
	CONTAINER_NAME=$(REGISTRY)/$(IMAGE_DEV) \
	docker-compose up -d
down: ## Remove application dev container
	@cd .docker && \
	APP_NAME=$(APP_NAME) \
	CONTAINER_NAME=$(REGISTRY)/$(IMAGE_DEV) \
	docker-compose down

console: ## Enter application dev container
	@docker exec -it -u developer $(APP_NAME) bash
console-root: ## Enter application dev container
	@docker exec -it -u root $(APP_NAME) bash

## Inside container

go-build: ## Build dev application (go build)	
	@env CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} GOARM=${GOARM} go build -o bin/${APP_NAME} ./

clean: ## Clean bin/
	@rm -rf bin/${APP_NAME}

remote-init: ## Init remote RPI for christmastree (connection, service...)
	if ! [ -f ~/.ssh/id_rsa.pub ]; then ssh-keygen; fi
	ssh-copy-id $(RPI_USER)@$(RPI_IP)
	ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -d ~/bin ]; then mkdir bin; fi'
	# scp ./bin/$(APP_NAME) $(RPI_USER)@$(RPI_IP):~/bin/
	scp ./raspbian/christmastree.service $(RPI_USER)@$(RPI_IP):~/bin/
	ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -f /usr/bin/christmastree ]; then sudo ln -s /home/$(RPI_USER)/bin/christmastree /usr/bin/; fi'
	ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -f /etc/systemd/system/christmastree.service ]; then sudo ln -s /home/$(RPI_USER)/bin/christmastree.service /etc/systemd/system/; fi'
	ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl enable christmastree.service'

remote-install: ## Install on remote RPI (send binary to RPI and restart service)
	scp ./bin/$(APP_NAME) $(RPI_USER)@$(RPI_IP):~/bin/
	ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl restart christmastree.service'

remote-start: ## Start christmastree service on RPI
	ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl start christmastree.service'

remote-stop: ## Stop christmastree service on RPI
	ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl stop christmastree.service'

