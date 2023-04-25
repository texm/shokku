#!/bin/bash
set -e

# This script installs and sets up shokku with a few prerequisites:
# dokku must be installed
# a global domain must be set
# letsencrypt must be configured and enabled with global email set

DOKKU_STORAGE_DIR="/var/lib/dokku/data/storage/"
SHOKKU_DATA_DIR="$DOKKU_STORAGE_DIR/shokku"
SHOKKU_APP_DATA_MOUNT_PATH="$SHOKKU_DATA_DIR:/data"
SHOKKU_VERSION=${SHOKKU_VERSION:-"latest"}
SHOKKU_IMAGE="ghcr.io/texm/shokku:$SHOKKU_VERSION"

SHOKKU_DOKKU_USER="shokkuadmin"
DISTROLESS_NONROOT_UID="65532"

clean-shokku() {
  echo "=> checking for existing resources"

  if dokku apps:exists shokku &> /dev/null; then
    echo "==> destroying old dokku app"
    dokku apps:destroy --force shokku &> /dev/null
  fi

  if dokku ssh-keys:list "$SHOKKU_DOKKU_USER" &> /dev/null; then
    echo "==> removing existing dokku ssh key"
    dokku ssh-keys:remove $SHOKKU_DOKKU_USER;
  fi

  echo "==> done"
}

create-shokku-app() {
  clean-shokku

  echo "=> pulling image (version: $SHOKKU_VERSION)"
  HOST_SSH_PORT=$(grep "Port " /etc/ssh/sshd_config | awk '{ print $2 }')
  docker pull "$SHOKKU_IMAGE" &> /dev/null
  SHOKKU_IMAGE_DIGEST=$(docker inspect --format='{{index .RepoDigests 0}}' "$SHOKKU_IMAGE")

  echo "=> creating & configuring dokku app"
  dokku apps:create shokku &> /dev/null
  dokku docker-options:add shokku deploy \
    "--add-host=host.docker.internal:host-gateway" &> /dev/null
  dokku config:set shokku \
    DOKKU_SSH_HOST='host.docker.internal' \
    DOKKU_SSH_PORT="$HOST_SSH_PORT" &> /dev/null

  echo "==> creating storage"
  dokku storage:ensure-directory shokku --chown false &> /dev/null
  dokku storage:mount shokku "$SHOKKU_APP_DATA_MOUNT_PATH" &> /dev/null
  chown -R "$DISTROLESS_NONROOT_UID":"$DISTROLESS_NONROOT_UID" "$SHOKKU_DATA_DIR" &> /dev/null

  echo "==> bootstrapping"
  dokku config:set shokku DOKKU_SKIP_DEPLOY=true &> /dev/null
  dokku git:from-image shokku "$SHOKKU_IMAGE_DIGEST" &> /dev/null

  shokku_ssh_key=$(dokku run shokku bootstrap) &> /dev/null
  echo "$shokku_ssh_key" | dokku ssh-keys:add "$SHOKKU_DOKKU_USER" &> /dev/null

  echo "==> deploying"
  dokku config:unset shokku DOKKU_SKIP_DEPLOY &> /dev/null

  echo "==> enabling letsencrypt"
  dokku letsencrypt:enable shokku &> /dev/null
}

main() {
  if [[ "$(id -u)" != "0" ]]; then
     echo "This script must be run as root" 1>&2
     exit 1
  fi

  if ! command -v dokku &> /dev/null; then
      echo "Please install dokku first using the instructions at https://dokku.com" 1>&2
      exit
  fi

  if ! dokku plugin:installed letsencrypt; then
    echo "Please setup letsencrypt using the instructions at https://dokku.com/docs/deployment/application-deployment/#setting-up-ssl" 1>&2
    exit
  fi

  for plugin in redis postgres mongo mysql; do
    if ! dokku plugin:installed $plugin; then
      echo "=> Installing plugin $plugin"
      dokku plugin:install https://github.com/dokku/dokku-$plugin.git $plugin &> /dev/null
    fi
  done

  create-shokku-app

  shokku_app_domain=$(dokku domains:report shokku --domains-app-vhosts)
  shokku_setup_key=$(dokku logs -q shokku | grep setup_key | jq ".setup_key")

  echo "=> shokku installed and running "
  echo "--- proceed with setup using key $shokku_setup_key at https://$shokku_app_domain ---"
}

main
