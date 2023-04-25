set shell := ["bash", "-uc"]

ui_npm := 'npm --prefix ui'
web_npm := 'npm --prefix web'
shokku_image := 'texm/shokku:latest'
dokku_image := 'dokku/dokku:0.30.2'
dokku_container_name := 'shokku-dokku-dev'
dokku_data_mount_dir := '/tmp/var/lib/dokku'
dokku_ssh_host := '127.0.0.1'
dokku_ssh_port := '3022'
db_path := 'test.db'

host_keyfile := ".test_keyfile"
container_keyfile := '/home/dokku/.ssh/authorized_keys'

reflex_script := "
# Run dev server
-s -r '\\.go$' -R '^ui/' -R '^\\.git$' -- just dev-backend

# Run ui dev server
# Exclude everything since vite will hot reload
-s -R '^.*' -- just dev-ui"

_default:
  @just --list

###
# Public recipes
###

# clean up resources from the backend and docker
@clean: _clean-backend _clean-docker

# setup dependencies, bootstrap backend, create dokku container
@setup: clean _install-dependencies _setup-dokku _setup-backend

# run the development environment (ui and backend)
@dev:
	-echo -n '{{reflex_script}}' | reflex -d 'fancy' -c '-'

@dev-backend: _touch-cmd-dist _setup-dokku
	just _run-with-env go run ./cmd/shokku

@dev-ui:
	-{{ui_npm}} run dev

# build a static /cmd/shokku binary
@build:
	{{ui_npm}} run build
	mv ./ui/dist ./cmd/shokku
	# GOOS=linux GOARCH=amd64 go build -o shokku ./cmd/shokku
	-go build -o shokku ./cmd/shokku
	-rm -r ./cmd/shokku/dist

@format:
	go fmt ./...
	{{ui_npm}} run prettier

# run the website development server
@dev-web:
	{{web_npm}} install
	{{web_npm}} run dev

# build the docker image and tag as latest
build-docker:
	docker build -t "{{shokku_image}}" .

# run the docker image 'texm/shokku:latest' with environment vars set
run-docker:
	@docker run -d \
		-e DOKKU_SSH_HOST='{{dokku_ssh_host}}' \
		-e DOKKU_SSH_PORT='{{dokku_ssh_port}}' \
		-e DB_PATH='{{db_path}}' \
		{{shokku_image}}

###
# Private helper recipes
###

## dependencies
@_install-dependencies:
	go install github.com/cespare/reflex@latest
	{{ui_npm}} install > /dev/null
	go mod tidy
##

## helpers
@_touch-cmd-dist:
	mkdir -p ./cmd/shokku/dist && touch ./cmd/shokku/dist/bleh

@_run-with-env +CMD:
	DOKKU_SSH_HOST='{{dokku_ssh_host}}' \
	DOKKU_SSH_PORT='{{dokku_ssh_port}}' \
	DB_PATH='{{db_path}}' \
	DEBUG_MODE=true \
		{{CMD}}
##

## builds & cleaning
@_clean-backend:
	go clean
	-rm -r ./cmd/shokku/dist &> /dev/null
	-rm {{db_path}} &> /dev/null
	-rm {{host_keyfile}} &> /dev/null

@_clean-docker:
	-docker rm -f {{dokku_container_name}} &> /dev/null
	-sudo rm -r {{dokku_data_mount_dir}} &> /dev/null
##

## Setup

@_setup-backend: _touch-cmd-dist
	just _run-with-env go run ./cmd/shokku bootstrap > {{host_keyfile}}
	just _add-dokku-ssh-key;

_setup-dokku:
	#!/bin/bash
	if [[ -z $(docker ps -a | grep '{{dokku_container_name}}') ]]; then
		just _run-dokku-container;
		just _install-dokku-plugins;
	elif [[ `docker container inspect -f '{{{{.State.Running}}' '{{dokku_container_name}}'` == 'false' ]]; then
		docker start {{dokku_container_name}};
	fi

@_run-dokku-container:
	docker run -d \
		--env DOKKU_HOSTNAME=dokku.me \
		--env DOKKU_HOST_ROOT={{dokku_data_mount_dir}}/home/dokku \
		--env DOKKU_LIB_HOST_ROOT={{dokku_data_mount_dir}}/var/lib/dokku \
		--name {{dokku_container_name}} \
		--publish {{dokku_ssh_port}}:22 \
	   	--volume {{dokku_data_mount_dir}}:/mnt/dokku \
		--volume /var/run/docker.sock:/var/run/docker.sock \
		{{dokku_image}} > /dev/null

	if [[ "{{ os() }}" == 'macos' ]]; then \
		docker exec shokku-dokku-dev bash -c \
			'groupmod -g 99 systemd-timesync && groupmod -g 101 docker'; fi

@_install-dokku-plugins:
	echo "installing dokku plugins"
	for plugin in redis postgres mongo mysql letsencrypt; do \
		echo "installing $plugin"; \
		docker exec {{dokku_container_name}} bash -c \
			"dokku plugin:install https://github.com/dokku/dokku-$plugin.git" \
			> /dev/null; \
	done

	# to fix plugins: https://github.com/dokku/dokku/issues/5004
	#	-docker exec {{dokku_container_name}} bash -c "rm -r /mnt/dokku/services" 2&> /dev/null
	#	-docker exec {{dokku_container_name}} bash -c "mv /var/lib/dokku/services /mnt/dokku/"
	#	-docker exec {{dokku_container_name}} bash -c "ln -s /mnt/dokku/services/ /var/lib/dokku/services"

@_add-dokku-ssh-key:
	docker exec {{dokku_container_name}} bash -c \
		'cd /home/ \
		 && rm -f {{container_keyfile}} \
		 && touch {{container_keyfile}} \
		 && chown dokku:dokku {{container_keyfile}}'

	docker exec {{dokku_container_name}} bash -c \
		"echo '$(cat {{host_keyfile}})' | dokku ssh-keys:add admin" > /dev/null
##