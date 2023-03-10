IMAGE_NAME = "revett/atlas-local-development:latest"

# Build the CLI binary locally.
build-cli:
	go build -o atlas

# Start a Docker container for local development.
contribute: docker-build-local-development-image
	docker run \
		--interactive \
		--name "atlas-local-development" \
		--rm \
		--tty \
		--volume $(shell pwd):/root/atlas \
		$(IMAGE_NAME) bash

# Build the Docker image for local development.
# https://github.com/docker/scan-cli-plugin/issues/149
docker-build-local-development-image:
	DOCKER_SCAN_SUGGEST=false docker build \
		--no-cache \
		--tag $(IMAGE_NAME) \
		./docker

# Install the CLI locally.
install-cli-locally: build
	rm -f ~/go/bin/atlas
	cp atlas ~/go/bin/
