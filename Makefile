include .env

APP_NAME = christmastree

.DEFAULT_GOAL = go-build

#-----------------------------------------------------------------------------------------------------------------------
ARG := $(word 2, $(MAKECMDGOALS))
%:
	@:
#-----------------------------------------------------------------------------------------------------------------------
#-----------------------------------------------------------------------------------------------------------------------

help: ## Outputs this help screen
	@grep -hE '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'


init: ## Init environment (add arm support && build ws2811-builder)
	@docker run --rm --privileged docker/binfmt:a7996909642ee92942dcd6cff44b9b95f08dad64
	@docker buildx build --platform linux/arm/v7 --tag ws2811-builder --file .docker/images/app-builder/Dockerfile .

go-modtidy: # Run go mod tidy
	@echo 'go mod tidy...'
	@docker run --rm \
		-v "$(PWD)":/usr/src/$(APP_NAME) \
		-v "$(PWD)/var/go:/go" \
		-v "$(PWD)/var/cache:/root/.cache" \
		--name ws2811-builder \
		--platform linux/arm/v7 \
  		-w /usr/src/$(APP_NAME) \
		ws2811-builder:latest \
		go mod tidy

go-build: ## Run go build
	@echo 'go build -o "bin/$(APP_NAME)"...'
	@docker run --rm \
		-v "$(PWD)":/usr/src/$(APP_NAME) \
		-v "$(PWD)/var/go:/go" \
		-v "$(PWD)/var/cache:/root/.cache" \
		--name ws2811-builder \
		--platform linux/arm/v7 \
  		-w /usr/src/$(APP_NAME) \
		ws2811-builder:latest \
		env CGO_ENABLED=1 go build -o "bin/$(APP_NAME)" -v

clean: ## Remove binary from bin/ directory
	@rm -rf bin/${APP_NAME}

## RPI commands
rpi-uptime: ## Get uptime from RPI
	@echo "RPI $(RPI_IP) uptime..."
	@ssh $(RPI_USER)@$(RPI_IP) 'uptime'

rpi-authorize: ## (keygen &&) ssh-copy-id
	@echo "RPI $(RPI_IP) authorize... (ssh-keygen, ssh-copy-id)"
	@if ! [ -f ~/.ssh/id_rsa.pub ]; then echo "ssh-keygen" && ssh-keygen; fi
	@ssh-copy-id $(RPI_USER)@$(RPI_IP) -f

rpi-install: ## Send binary and config to RPI
	@echo "Send binary and config to RPI"
	scp ./bin/$(APP_NAME) $(RPI_USER)@$(RPI_IP):~/bin/
	@ssh $(RPI_USER)@$(RPI_IP) 'rm ~/bin/config/0*.yaml -f'
	scp ./config/* $(RPI_USER)@$(RPI_IP):~/bin/config/


rpi-enable-service: ## Enable christmastree service on RPI
	@echo "Enable christmastree service on RPI $(RPI_IP)..."
	@ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -d ~/bin ]; then mkdir ~/bin; fi'
	@ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -d ~/bin/config ]; then mkdir ~/bin/config; fi'
	@scp ./raspbian/christmastree.service $(RPI_USER)@$(RPI_IP):~/bin/
	ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -d /etc/$(APP_NAME) ]; then sudo ln -s ~/bin/config /etc/$(APP_NAME); fi'
	@ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -f /usr/bin/christmastree ]; then sudo ln -s /home/$(RPI_USER)/bin/christmastree /usr/bin/; fi'
	@ssh $(RPI_USER)@$(RPI_IP) 'if ! [ -f /etc/systemd/system/christmastree.service ]; then sudo ln -s /home/$(RPI_USER)/bin/christmastree.service /etc/systemd/system/; fi'
	@ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl enable christmastree.service'

rpi-start-service: ## Start christmastree service on RPI
	@echo "Start christmastree service on RPI $(RPI_IP)..."
	@ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl start christmastree.service'

rpi-stop-service: ## Stop christmastree service on RPI
	@echo "Start christmastree service on RPI $(RPI_IP)..."
	@ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl stop christmastree.service'

rpi-restart-service: ## Restart christmastree service on RPI
	@echo "Restart christmastree service on RPI $(RPI_IP)..."
	@ssh $(RPI_USER)@$(RPI_IP) 'sudo systemctl restart christmastree.service'

rpi-down: ## Poweroff RPI
	@echo "Send 'poweroff' to RPI $(RPI_IP)..."
	@ssh $(RPI_USER)@$(RPI_IP) 'sudo poweroff'

rpi-console: ## SSH RPI console
	@ssh $(RPI_USER)@$(RPI_IP)